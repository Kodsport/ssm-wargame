package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"

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

	user_service "github.com/sakerhetsm/ssm-wargame/internal/api/user"
	user_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/user/server"
	user_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/user"

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
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer log.Sync()

	mux := goahttp.NewMuxer()

	auther := auth.New(cfg, log, db)

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
		svc := challenge_service.NewService(db, log, auther, s3c)
		endpoints := challenge_transport.NewEndpoints(svc)
		s := challenge_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		s.Use(RequestInContext())
		challenge_server.Mount(mux, s)
	}
	{
		svc := auth_service.NewService(db, log, cfg)
		endpoints := auth_transport.NewEndpoints(svc)
		s := auth_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		auth_server.Mount(mux, s)
	}
	{
		svc := admin_service.NewService(db, log, auther, s3c, cfg)
		endpoints := admin_transport.NewEndpoints(svc)
		s := admin_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		admin_server.Mount(mux, s)
	}
	{
		svc := user_service.NewService(db, log, auther)
		endpoints := user_transport.NewEndpoints(svc)
		s := user_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		s.Use(goahttpmid.RequestID())
		user_server.Mount(mux, s)
	}

	var handler http.Handler = mux

	// oops this is dev only
	// Uncomment when cors needs to be modified, ex. when api is on localhost:8000 and app is on localhost:3000
	// handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("access-control-allow-origin", `*`)
	// 	w.Header().Add("access-control-allow-headers", r.Header.Get("Access-Control-Request-Headers"))
	//
	// 	if r.Method == http.MethodOptions {
	// 		w.WriteHeader(http.StatusNoContent)
	// 		return
	// 	}
	//
	// 	mux.ServeHTTP(w, r)
	// })
	//
	//	handler = cors.New(cors.Options{
	//		AllowedOrigins: []string{"sakerhetssm.se"},
	//		MaxAge:         60 * 60 * 24,
	//	}).Handler(handler)

	srv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: handler,
	}

	log.Info("starting http server")
	return srv.ListenAndServe()
}

func RequestInContext() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "ssm_http_req", r)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
