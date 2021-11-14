package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/sakerhetsm/ssm-wargame/internal/api"
	"github.com/sakerhetsm/ssm-wargame/internal/db"

	challenge_transport "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	challenge_server "github.com/sakerhetsm/ssm-wargame/internal/gen/http/challenge/server"

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

	svc := api.NewService(db.New(conn))
	endpoints := challenge_transport.NewEndpoints(svc)
	s := challenge_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	challenge_server.Mount(mux, s)

	var handler http.Handler = mux

	srv := &http.Server{Addr: "localhost:8080", Handler: handler}

	return srv.ListenAndServe()
}
