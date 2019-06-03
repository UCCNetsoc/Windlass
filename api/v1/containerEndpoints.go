package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/UCCNetworkingSociety/Windlass/middleware"

	common "github.com/UCCNetworkingSociety/Windlass/api/models"

	"github.com/UCCNetworkingSociety/Windlass/app/services"

	"github.com/go-chi/chi"
)

type ContainerEndpoint struct{}

// TODO sort out endpoint URIs
func NewContainerEndpoints(r chi.Router) {
	e := ContainerEndpoint{}
	r.Get("/container", middleware.WithContext(e.createContainer, time.Second*3))
}

func (e ContainerEndpoint) createContainer(w http.ResponseWriter, r *http.Request) {
	err := services.NewContainerHostService().
		WithContext(r.Context()).
		CreateHost("sample text")
	if err != nil {
		json.NewEncoder(w).Encode(common.APIResponse{
			Status:  http.StatusInternalServerError,
			Content: "error creating container",
		})
		return
	}
}
