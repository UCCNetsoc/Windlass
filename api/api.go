package api

import (
	v1 "github.com/UCCNetworkingSociety/Windlass/api/v1"
	"github.com/go-chi/chi"
)

type API struct {
	routes *chi.Mux
}

func NewAPI(router *chi.Mux) *API {
	return &API{
		routes: router,
	}
}

func (api *API) Init() {
	api.routes.Route("/v1", func(r chi.Router) {
		v1.NewContainerEndpoints(r)
	})
}
