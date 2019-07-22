package api

import (
	v1 "github.com/UCCNetworkingSociety/Windlass/app/api/v1"
	"github.com/UCCNetworkingSociety/Windlass/middleware"
	"github.com/go-chi/chi"
	middlechi "github.com/go-chi/chi/middleware"
)

type API struct {
	routes chi.Router
}

func NewAPI(router *chi.Mux) *API {
	return &API{
		routes: router,
	}
}

func (api *API) Init() {
	api.routes.Use(middlechi.RealIP)
	api.routes.Use(middleware.Recoverer)
	api.routes.Route("/v1", func(r chi.Router) {
		v1.NewProjectEndpoints(r)
	})
}
