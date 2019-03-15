package api

import (
	"encoding/json"
	"github.com/UCCNetworkingSociety/Windlass/types"
	"github.com/fsouza/go-dockerclient"
	"net/http"
)

func (a *API) listContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := a.servers.Docker.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(types.APIError{
			ErrorID: 100,
			APIResponse: types.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			},
		})
		return
	}
	json.NewEncoder(w).Encode(containers)
}
