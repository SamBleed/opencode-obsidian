package server

import (
	"fmt"
	"github.com/SamBleed/ozy-api/pkg/logger"
	"net/http"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{\"status\": \"healthy\"}")
	})

	logger.Info("Starting server", map[string]interface{}{"port": s.port})
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), mux)
}
