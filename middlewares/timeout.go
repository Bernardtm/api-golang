package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"
)

// TimeoutMiddleware is a middleware that sets a timeout for all requests
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Creates a new context with a timeout
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			// Replaces the request with the new context
			r = r.WithContext(ctx)

			// Channel to capture if the handler has completed
			done := make(chan struct{})

			go func() {
				// Executes the next handler
				next.ServeHTTP(w, r)
				close(done)
			}()

			select {
			case <-done:
				// If the handler completed within the timeout
				return
			case <-ctx.Done():
				// If the timeout is reached, cancel the request and respond with an error
				log.Printf("Request timeout: %s %s", r.Method, r.URL.Path)
				http.Error(w, "Request timeout", http.StatusGatewayTimeout)
			}
		})
	}
}
