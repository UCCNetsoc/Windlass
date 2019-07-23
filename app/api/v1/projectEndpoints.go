package v1

import (
	"net/http"
	"time"

	midware "github.com/UCCNetworkingSociety/Windlass/middleware"

	"github.com/UCCNetworkingSociety/Windlass/app/api/models"
	"github.com/go-chi/render"

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
		r.Post("/", midware.WithContext(p.create, time.Second*15))
	})
}

func (p *ProjectEndpoint) create(w http.ResponseWriter, r *http.Request) {
	var newProject project.Project

	if err := render.Bind(r, &newProject); err != nil {
		render.Render(w, r, models.APIResponse{
			Status:  http.StatusBadRequest,
			Content: err.Error(),
		})
		return
	}

	if err := p.projectService.CreateProject(r.Context(), newProject); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err).Error("error occurred")
	}
}
