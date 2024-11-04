package main

import (
	"context"
	"fmt"
	common "github.com/ErikPelli/requests_concurrency_benchmark"
	"net/http"
	"strconv"
	"time"
)

func printHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type server struct {
	server http.Server
}

func newEchoServer(port int) *server {
	return &server{
		server: http.Server{
			Addr:         ":" + strconv.Itoa(port),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      http.HandlerFunc(printHandler),
		},
	}
}

func (s *server) Start() error {
	return s.server.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func main() {
	s := newEchoServer(common.Port)
	fmt.Println(s.Start())
	s.Stop(context.Background())
}
