package main

import (
	"net/http"

	"github.com/Strum355/viper"

	"github.com/UCCNetworkingSociety/Windlass/api"
	"github.com/UCCNetworkingSociety/Windlass/app/config"
	"github.com/UCCNetworkingSociety/Windlass/app/connections"
	"github.com/UCCNetworkingSociety/Windlass/utils/logging"
	"github.com/UCCNetworkingSociety/Windlass/utils/must"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	v := viper.GetViper()
	config.Load(v)

	must.Do(connections.EstablishConnections)
	defer connections.Group.Close()

	api.NewAPI(r).Init()
	log.Info("API server started")

	http.ListenAndServe(":8080", r)
}
