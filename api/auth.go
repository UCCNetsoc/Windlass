package api

import (
	"net/http"
)

type Token struct {

}

func (a *API) CheckAPIToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}