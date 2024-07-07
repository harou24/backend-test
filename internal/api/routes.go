package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(router *mux.Router, breedHandler *BreedHandler) {
	router.HandleFunc("/breeds/search", breedHandler.SearchBreedsHandler).Methods(http.MethodGet)
	router.HandleFunc("/breeds", breedHandler.GetBreedsHandler).Methods(http.MethodGet)
	router.HandleFunc("/breeds", breedHandler.CreateBreedHandler).Methods(http.MethodPost)
	router.HandleFunc("/breeds/{id}", breedHandler.GetBreedHandler).Methods(http.MethodGet)
	router.HandleFunc("/breeds/{id}", breedHandler.UpdateBreedHandler).Methods(http.MethodPut)
	router.HandleFunc("/breeds/{id}", breedHandler.DeleteBreedHandler).Methods(http.MethodDelete)
}
