package main

import (
	"github.com/UCCNetworkingSociety/Windlass/api"
	"log"
	"github.com/UCCNetworkingSociety/Windlass/types"
	"net/http"
	"github.com/go-chi/chi"
)

func main() { 
	r := chi.NewRouter()
	s, err := types.NewServerGroup()
	if err != nil {
		log.Fatalln(err)
	}
	defer s.Close()
	
	api.NewAPI(s, r).SetupRoutes()

	http.ListenAndServe(":8080", r)
}