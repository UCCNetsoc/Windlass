package api

import (
	"github.com/go-chi/chi"
	"github.com/UCCNetworkingSociety/CaaS-Backend/server"
)

type API struct {
	s *server.ServerGroup
	Routes *chi.Mux
}