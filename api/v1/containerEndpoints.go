package v1

import (
	"encoding/json"
	"net/http"

	"github.com/UCCNetworkingSociety/Windlass/types"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/go-chi/chi"
)

type ContainerEndpoint struct {
	servers *types.ServerGroup
}

func NewContainerEndpoints(r chi.Router, group *types.ServerGroup) {
	e := ContainerEndpoint{group}
	r.Get("/containers", e.listContainers)
}

func (e ContainerEndpoint) listContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := e.servers.Docker.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(types.APIError{
			ErrorID: 100,
			APIResponse: types.APIResponse{
				Status:  http.StatusInternalServerError,
				Content: err.Error(),
			},
		})
		return
	}
	json.NewEncoder(w).Encode(containers)
}
