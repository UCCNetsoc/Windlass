package api

import (
	"net/http"

	"github.com/UCCNetworkingSociety/Windlass/app/api/models"
	v1 "github.com/UCCNetworkingSociety/Windlass/app/api/v1"
	"github.com/UCCNetworkingSociety/Windlass/middleware"
	"github.com/go-chi/chi"
	middlechi "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type API struct {
	routes chi.Router
}

func NewAPI(router *chi.Mux) *API {
	return &API{
		routes: router,
	}
}

func (api *API) Init() {
	api.routes.Use(middlechi.RealIP)
	api.routes.Use(middlechi.DefaultLogger)
	api.routes.Use(middleware.Recoverer)
	api.routes.Use(render.SetContentType(render.ContentTypeJSON))

	api.routes.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &models.APIResponse{
			Status: http.StatusOK,
		})
	})
	api.routes.Route("/v1", func(r chi.Router) {
		v1.NewProjectEndpoints(r)
	})
}
