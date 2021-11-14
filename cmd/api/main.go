package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sakerhetsm/ssm-wargame/internal/api"

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

	mux := goahttp.NewMuxer()

	svc := api.NewService()
	endpoints := challenge_transport.NewEndpoints(svc)
	s := challenge_server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	challenge_server.Mount(mux, s)

	var handler http.Handler = mux

	srv := &http.Server{Addr: "localhost:8080", Handler: handler}

	return srv.ListenAndServe()
}
