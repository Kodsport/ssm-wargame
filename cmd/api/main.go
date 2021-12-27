package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/sakerhetsm/ssm-wargame/internal/auth"
	"github.com/sakerhetsm/ssm-wargame/internal/config"

	challenge_service "github.com/sakerhetsm/ssm-wargame/internal/api/challenge"
	challenge_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challenge_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/challenge/server"

	auth_service "github.com/sakerhetsm/ssm-wargame/internal/api/auth"
	auth_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	auth_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/auth/server"

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
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer log.Sync()

	mux := goahttp.NewMuxer()

	auther := auth.New(cfg, log, conn)

	{
		svc := challenge_service.NewService(conn, log, auther)
		endpoints := challenge_transport.NewEndpoints(svc)
		s := challenge_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		challenge_server.Mount(mux, s)
	}
	{
		svc := auth_service.NewService(conn, log, cfg)
		endpoints := auth_transport.NewEndpoints(svc)
		s := auth_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		auth_server.Mount(mux, s)
	}

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Info("starting http server")
	return srv.ListenAndServe()
}
