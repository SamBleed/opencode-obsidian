package server

import (
	"github.com/SamBleed/ozy-api/pkg/logger"
	"net/http"
	"time"
)

// LoggingMiddleware registra cada petición en formato JSON
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Aquí podríamos generar o leer un Trace-ID del header
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = "gen-" + string(time.Now().UnixNano())
		}

		next.ServeHTTP(w, r)

		logger.Info("Request processed", map[string]interface{}{
			"method":    r.Method,
			"path":      r.URL.Path,
			"duration":  time.Since(start).String(),
			"remote_ip": r.RemoteAddr,
			"trace_id":  traceID,
		})
	})
}
