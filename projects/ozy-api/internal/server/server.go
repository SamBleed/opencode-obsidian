package server

import (
	"encoding/json"
	"fmt"
	"github.com/SamBleed/ozy-api/internal/domain"
	"github.com/SamBleed/ozy-api/internal/repository/auth"
	"github.com/SamBleed/ozy-api/internal/repository/db"
	"github.com/SamBleed/ozy-api/pkg/logger"
	"net/http"
	"strings"
)

type Server struct {
	port        int
	authAdapter *auth.JWTAdapter
	entityRepo  domain.EntityRepository
}

func NewServer(port int, secret string, dbAdapter *db.PostgresAdapter) *Server {
	return &Server{
		port:        port,
		authAdapter: auth.NewJWTAdapter(secret),
		entityRepo:  db.NewSqlEntityRepository(dbAdapter),
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Públicos
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/login", s.handleLogin)

	// Entidades (Protegidos)
	mux.HandleFunc("/api/entities", s.AuthMiddleware(s.handleEntities))

	handler := LoggingMiddleware(mux)

	logger.Info("Starting server: Entity Management Active", map[string]interface{}{"port": s.port})
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), handler)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{\"status\": \"healthy\"}")
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	token, _ := s.authAdapter.GenerateToken("sam-123")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (s *Server) handleEntities(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		entities, err := s.entityRepo.List()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)

	case http.MethodPost:
		var e domain.Entity
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if err := s.entityRepo.Save(&e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "{\"message\": \"Entity saved successfully\"}")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
