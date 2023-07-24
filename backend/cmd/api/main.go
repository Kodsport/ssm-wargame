package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/config"

	challenge_service "github.com/sakerhetsm/ssm-wargame/internal/api/challenge"
	challenge_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challenge_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/challenge/server"

	auth_service "github.com/sakerhetsm/ssm-wargame/internal/api/auth"
	auth_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	auth_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/auth/server"

	admin_service "github.com/sakerhetsm/ssm-wargame/internal/api/admin"
	admin_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	admin_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/admin/server"

	goahttp "goa.design/goa/v3/http"
	goahttpmid "goa.design/goa/v3/http/middleware"
)

func main() {
	err := realMain()
	if err != nil {
		fmt.Println("error in main:", err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	godotenv.Load()
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Address, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode)
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer log.Sync()

	mux := goahttp.NewMuxer()

	auther := auth.New(cfg, log, pool)

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         &cfg.S3.Endpoint,
		Credentials:      credentials.NewStaticCredentials(cfg.S3.KeyID, cfg.S3.KeySecret, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return err
	}

	s3c := s3.New(sess)

	{
		svc := challenge_service.NewService(pool, log, auther, s3c)
		endpoints := challenge_transport.NewEndpoints(svc)
		s := challenge_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		challenge_server.Mount(mux, s)
	}
	{
		svc := auth_service.NewService(pool, log, cfg)
		endpoints := auth_transport.NewEndpoints(svc)
		s := auth_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		auth_server.Mount(mux, s)
	}
	{
		svc := admin_service.NewService(pool, log, auther, s3c, cfg)
		endpoints := admin_transport.NewEndpoints(svc)
		s := admin_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		admin_server.Mount(mux, s)
	}

	var handler http.Handler = mux

	handler = cors.AllowAll().Handler(handler)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	log.Info("starting http server")
	return srv.ListenAndServe()
}
