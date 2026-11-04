package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string) *Server {
	s := &Server{
		srv: &http.Server{
			Addr: addr,
		},
	}
	s.srv.Handler = s.handler()
	return s
}

func serverStart() {

	addr := ":8080"
	srvhandler := NewServer(addr)
	srvhandler.srv.ListenAndServe()

}

func (s *Server) handler() *http.ServeMux {

	srvMux := http.NewServeMux()

	srvMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("content/type", "plain/text")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "hello world")
		}
	})

	return srvMux

}

func main() {
	serverStart()
}
