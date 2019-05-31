package v1

import (
	"net/http"

	"github.com/UCCNetworkingSociety/Windlass/app/services"

	"github.com/go-chi/chi"
)

type ContainerEndpoint struct{}

func NewContainerEndpoints(r chi.Router) {
	e := ContainerEndpoint{}
	r.Get("/containers", e.listContainers)
}

func (e ContainerEndpoint) listContainers(w http.ResponseWriter, r *http.Request) {
	services.NewContainerHostService().
		WithContext(r.Context()).
		CreateHost("sample_text")
}
