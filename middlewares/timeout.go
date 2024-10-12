package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"
)

// TimeoutMiddleware é um middleware que define um tempo limite para todas as requisições
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Cria um novo contexto com timeout
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			// Substitui o request com o novo contexto
			r = r.WithContext(ctx)

			// Canal para capturar se o handler foi completado
			done := make(chan struct{})

			go func() {
				// Executa o próximo handler
				next.ServeHTTP(w, r)
				close(done)
			}()

			select {
			case <-done:
				// Se o handler foi completado dentro do tempo limite
				return
			case <-ctx.Done():
				// Se o timeout foi atingido, cancela a requisição e responde com erro
				log.Printf("Request timeout: %s %s", r.Method, r.URL.Path)
				http.Error(w, "Request timeout", http.StatusGatewayTimeout)
			}
		})
	}
}
