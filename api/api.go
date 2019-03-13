package api

import (
	"net/http"
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

func (a *API) Routes() {
	r := chi.NewRouter()
	r.Get("/containers", a.listContainers)
}

func (a *API) listContainers(w http.ResponseWriter, r *http.Request) {

}