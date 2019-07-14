package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/UCCNetworkingSociety/Windlass/app/api/models"
	"github.com/UCCNetworkingSociety/Windlass/app/services"
	"github.com/UCCNetworkingSociety/Windlass/middleware"
	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
	"github.com/go-chi/chi"
)

type ContainerEndpoint struct{}

// TODO sort out endpoint URIs
func NewContainerEndpoints(r chi.Router) {
	e := ContainerEndpoint{}
	r.Get("/container", middleware.WithContext(e.createContainer, time.Second*20))
}

func (e ContainerEndpoint) createContainer(w http.ResponseWriter, r *http.Request) {
	containerHostName := r.URL.Query().Get("name")
	if containerHostName == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  http.StatusBadRequest,
			Content: "project name can't be empty",
		})
		return
	}
	err := services.NewContainerHostService().
		WithContext(r.Context()).
		CreateHost(containerHostName)
	if err != nil {
		log.Error("error creating container - %v", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  http.StatusInternalServerError,
			Content: "error creating project",
		})
		return
	}
}
