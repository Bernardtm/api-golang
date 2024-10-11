package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Logger is a middleware that logs the details of each request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a ResponseWriter to capture the status code
		lrw := &LoggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r) // Call the next handler

		// Log the request details
		log.Printf(
			"%s %s %d %s",
			r.Method,
			r.URL.Path,
			lrw.statusCode,
			time.Since(start),
		)
	})
}

// LoggingResponseWriter is a wrapper around http.ResponseWriter to capture status codes.
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code in the LoggingResponseWriter
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
