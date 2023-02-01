package httpserver

import (
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        handler,
	}
	return s.httpServer.ListenAndServe()
}

func WritePORT() string {
	PORT := "8080"
	args := os.Args
	if len(args) == 1 {
		return ":" + PORT
	} else if len(args) == 2 {
		PORT = ":" + args[1]
	}
	return PORT
}
