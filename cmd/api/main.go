package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/db"
	"go.uber.org/zap"

	challenge_service "github.com/sakerhetsm/ssm-wargame/internal/api/challenge"
	challenge_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challenge_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/challenge/server"

	auth_service "github.com/sakerhetsm/ssm-wargame/internal/api/auth"
	auth_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
	auth_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/auth/server"

	goahttp "goa.design/goa/v3/http"
)

func main() {
	err := realMain()
	if err != nil {
		fmt.Println("error in main:", err.Error())
		os.Exit(1)
	}
}

func realMain() error {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/ssm_wargame?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	mux := goahttp.NewMuxer()

	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	tx, _ := conn.BeginTx(nil, pgx.TxOptions{})
	db.New(conn).WithTx(tx)

	{
		svc := challenge_service.NewService(conn, log)
		endpoints := challenge_transport.NewEndpoints(svc)
		s := challenge_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		challenge_server.Mount(mux, s)
	}
	{
		svc := auth_service.NewService(conn, log)
		endpoints := auth_transport.NewEndpoints(svc)
		s := auth_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
		auth_server.Mount(mux, s)
	}

	var handler http.Handler = mux

	srv := &http.Server{Addr: "localhost:8080", Handler: handler}

	return srv.ListenAndServe()
}
