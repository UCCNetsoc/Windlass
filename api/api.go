package api

import (
	"github.com/go-chi/chi"
	"github.com/UCCNetworkingSociety/Windlass/types"
)

type API struct {
	servers *types.ServerGroup
	routes *chi.Mux
}

func NewAPI(serverGroup *types.ServerGroup, router *chi.Mux) *API {
	return &API{
		servers: serverGroup,
		routes: router,
	}
}

func (a *API) SetupRoutes() {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/containers", a.listContainers)
	})
	a.routes.Mount("/api", r)
}