package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"

	"github.com/UCCNetworkingSociety/Windlass-worker/app/models/project"

	"github.com/UCCNetworkingSociety/Windlass/app/services"
	"github.com/go-chi/chi"
)

type ProjectEndpoint struct {
	projectService *services.ProjectService
}

// TODO sort out endpoint URIs
func NewProjectEndpoints(r chi.Router) {
	p := ProjectEndpoint{projectService: services.NewProjectService()}

	r.Route("/project", func(r chi.Router) {
		r.Post("/", p.create)
	})
}

func (p *ProjectEndpoint) create(w http.ResponseWriter, r *http.Request) {
	var postProject project.Project

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&postProject); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err := p.projectService.CreateProject(context.Background(), project.Project{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err, "error occurred")
	}
}
