package api

import (
	v1 "github.com/UCCNetworkingSociety/Windlass/api/v1"
	"github.com/UCCNetworkingSociety/Windlass/types"
	"github.com/go-chi/chi"
)

type API struct {
	servers *types.ServerGroup
	routes  *chi.Mux
}

func NewAPI(serverGroup *types.ServerGroup, router *chi.Mux) *API {
	return &API{
		servers: serverGroup,
		routes:  router,
	}
}

func (api *API) Init() {
	api.routes.Route("/v1", func(r chi.Router) {
		v1.NewContainerEndpoints(r, api.servers)
	})
}
