package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ContainerEndpoint struct{}

func NewContainerEndpoints(r chi.Router) {
	e := ContainerEndpoint{}
	r.Get("/containers", e.listContainers)
}

func (e ContainerEndpoint) listContainers(w http.ResponseWriter, r *http.Request) {
	/* containers, err := connections.Group.Docker.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(common.APIError{
			ErrorID: 100,
			APIResponse: common.APIResponse{
				Status:  http.StatusInternalServerError,
				Content: err.Error(),
			},
		})
		return
	}
	json.NewEncoder(w).Encode(containers) */
}
