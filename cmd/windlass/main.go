package main

import (
	"net/http"

	"github.com/Strum355/log"
	"github.com/UCCNetworkingSociety/Windlass/app/api"
	"github.com/UCCNetworkingSociety/Windlass/app/config"
	"github.com/UCCNetworkingSociety/Windlass/app/connections"
	"github.com/UCCNetworkingSociety/Windlass/utils/must"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	log.InitSimpleLogger(&log.Config{
		LogLevel: log.LogDebug,
	})

	must.Do(config.Load)

	must.Do(connections.EstablishConnections)
	defer connections.Close()

	config.PrintSettings()

	api.NewAPI(r).Init()
	log.Info("API server started")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.WithError(err).Error("error starting server")
	}
}
