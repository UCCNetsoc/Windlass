package main

import (
	"github.com/spf13/viper"
	"net/http"

	"github.com/UCCNetworkingSociety/Windlass/api"
	"github.com/UCCNetworkingSociety/Windlass/must"
	"github.com/UCCNetworkingSociety/Windlass/types"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	must.Do(func() error {
		viper.GetViper()
		return nil
	})

	var s *types.ServerGroup
	must.Do(func() error {
		var err error
		s, err = types.NewServerGroup()
		return err
	})
	defer s.Close()

	api.NewAPI(s, r).Init()

	http.ListenAndServe(":8080", r)
}
