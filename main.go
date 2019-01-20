package main

import (
	"github.com/UCCNetworkingSociety/CaaS-Backend/api"
	"log"
	"github.com/UCCNetworkingSociety/CaaS-Backend/types"
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
	
	a := api.NewAPI(s, r)

	http.ListenAndServe("8080", r)
}