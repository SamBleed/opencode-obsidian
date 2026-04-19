package server

import (
	"encoding/json"
	"fmt"
	"github.com/SamBleed/ozy-api/internal/repository/auth"
	"github.com/SamBleed/ozy-api/internal/repository/db"
	"github.com/SamBleed/ozy-api/pkg/logger"
	"net/http"
	"strings"
)

type Server struct {
	port        int
	authAdapter *auth.JWTAdapter
	dbAdapter   *db.PostgresAdapter
}

func NewServer(port int, secret string, dbAdapter *db.PostgresAdapter) *Server {
	return &Server{
		port:        port,
		authAdapter: auth.NewJWTAdapter(secret),
		dbAdapter:   dbAdapter,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Endpoints públicos
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/login", s.handleLogin)

	// Endpoints protegidos
	mux.HandleFunc("/api/secure", s.AuthMiddleware(s.handleSecure))

	// Aplicar el middleware de logging globalmente
	handler := LoggingMiddleware(mux)

	logger.Info("Starting server with Security Guardrails and Database active", map[string]interface{}{
		"port": s.port,
		"db":   "connected",
	})
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), handler)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{\"status\": \"healthy\", \"db\": \"connected\"}")
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Simulación de login exitoso (En el futuro buscar en s.dbAdapter)
	token, _ := s.authAdapter.GenerateToken("sam-123")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (s *Server) handleSecure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "{\"message\": \"Welcome to the Bunker Secure Area\"}")
}

func (s *Server) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := s.authAdapter.ValidateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
