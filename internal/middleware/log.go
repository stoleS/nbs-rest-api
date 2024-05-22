package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type withCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *withCodeWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		scheme := "http"

		withCode := &withCodeWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(withCode, r)

		requestLogger := log.WithFields(log.Fields{
			"Status Code": withCode.statusCode,
			"Method":      r.Method,
			"Path":        r.URL.Path,
			"Time":        time.Since(start),
		})

		log.SetFormatter(&log.TextFormatter{
			ForceQuote: true,
		})

		requestLogger.Logf(log.InfoLevel, "%s://%s%s", scheme, r.Host, r.RequestURI)
	})
}
