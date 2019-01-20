package api

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/UCCNetworkingSociety/CaaS-Backend/types"
)

type API struct {
	servers *types.ServerGroup
	routes *chi.Mux
}

func NewAPI(serverGroup *types.ServerGroup, router *chi.Mux) *API {
	return &API{
		s: serverGroup,
		r: router,
	}
}

func (a *API) Routes() {
	r := chi.NewRouter()
	r.Get("/containers", a.listContainers)
}

func (a *API) listContainers(w http.ResponseWriter, r *http.Request) {

}