package main

import (
	"net/http"

	"github.com/Strum355/viper"

	"github.com/UCCNetworkingSociety/Windlass/api"
	"github.com/UCCNetworkingSociety/Windlass/config"
	"github.com/UCCNetworkingSociety/Windlass/connections"
	"github.com/UCCNetworkingSociety/Windlass/logging"
	"github.com/UCCNetworkingSociety/Windlass/must"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	v := viper.GetViper()
	config.Load(v)

	must.Do(connections.EstablishConnections)
	defer connections.Group.Close()
	log.Info("connections established")

	api.NewAPI(r).Init()
	log.Info("API server started")

	http.ListenAndServe(":8080", r)
}
