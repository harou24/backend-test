package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(router *mux.Router, breedHandler *BreedHandler) {
	// Middleware to enable CORS for all routes
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/breeds/search", breedHandler.SearchBreedsHandler).Methods(http.MethodGet)
	router.HandleFunc("/breeds", breedHandler.GetBreedsHandler).Methods(http.MethodGet)

	// Handle POST method with CORS
	router.HandleFunc("/breeds", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		breedHandler.CreateBreedHandler(w, r)
	}).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/breeds/{id}", breedHandler.GetBreedHandler).Methods(http.MethodGet)

	// Handle PUT method with CORS
	router.HandleFunc("/breeds/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		breedHandler.UpdateBreedHandler(w, r)
	}).Methods(http.MethodPut, http.MethodOptions)

	// Handle DELETE method with CORS
	router.HandleFunc("/breeds/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		breedHandler.DeleteBreedHandler(w, r)
	}).Methods(http.MethodDelete, http.MethodOptions)
}
