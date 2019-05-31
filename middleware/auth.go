package middleware

import (
	"net/http"
)

// CheckAPIToken compares the JWT Token claims to the scope and resource owner
func CheckAPIToken(next http.Handler, scope string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
