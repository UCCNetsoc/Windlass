package middleware

import (
	"context"
	"net/http"
	"time"
)

// WithContext allows you to add a context timeout to a request
func WithContext(next http.HandlerFunc, timeout time.Duration) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
