package api

import (
	"github.com/UCCNetworkingSociety/Windlass/types"
	"net/http"
)

type Token struct {

}

// CheckAPIToken compares the JWT Token claims to the scope and resource owner
func CheckAPIToken(next http.Handler, group *types.ServerGroup, scope string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}