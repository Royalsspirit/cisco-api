package api

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// addRequestIDMiddleware add x-request-id in response header
func (s *Server) addRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid, err := uuid.NewV4()

		if err == nil {
			// Add x-request-id in response's header
			w.Header().Set("X-Request-Id", uuid.String())
		}

		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware log information information about the current request
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		s.Logger.WithFields(logrus.Fields{
			"method":       r.Method,
			"uri":          r.URL.RequestURI,
			"ip":           r.RemoteAddr,
			"duration":     time.Since(start),
			"x-request-id": w.Header().Get("X-Request-Id"),
		}).Info("request")
	})
}
