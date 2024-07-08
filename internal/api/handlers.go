package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/internal/domain"
	"net/http"
	"strconv"
)

type BreedHandler struct {
	breedRepository domain.BreedRepository
}

func NewBreedHandler(breedRepository domain.BreedRepository) *BreedHandler {
	return &BreedHandler{breedRepository: breedRepository}
}

func (h *BreedHandler) GetBreedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid breed ID", http.StatusBadRequest)
		return
	}

	breed, err := h.breedRepository.GetBreedByID(id)
	if err != nil {
		if err == domain.ErrBreedNotFound {
			http.Error(w, "Breed not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breed)
}

func (h *BreedHandler) GetBreedsHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	breeds, totalCount, err := h.breedRepository.GetAllBreeds(page, limit)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Total-Count", strconv.Itoa(totalCount))
	json.NewEncoder(w).Encode(breeds)
}

func (h *BreedHandler) CreateBreedHandler(w http.ResponseWriter, r *http.Request) {
	var breed domain.Breed
	err := json.NewDecoder(r.Body).Decode(&breed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBreed, err := h.breedRepository.CreateBreed(breed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBreed)
}

func (h *BreedHandler) UpdateBreedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid breed ID", http.StatusBadRequest)
		return
	}

	var breed domain.Breed
	err = json.NewDecoder(r.Body).Decode(&breed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.breedRepository.UpdateBreed(id, breed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedBreed, err := h.breedRepository.GetBreedByID(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBreed)
}

func (h *BreedHandler) DeleteBreedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid breed ID", http.StatusBadRequest)
		return
	}

	err = h.breedRepository.DeleteBreed(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BreedHandler) SearchBreedsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	species := queryParams.Get("species")
	weightStr := queryParams.Get("weight")

	criteria := make(map[string]any)

	if species != "" {
		criteria["species"] = species
	}

	if weightStr != "" {
		weightInt, err := strconv.Atoi(weightStr)
		if err != nil {
			http.Error(w, "Invalid weight parameter", http.StatusBadRequest)
			return
		}
		criteria["weight"] = weightInt
	}

	breeds, err := h.breedRepository.SearchBreeds(criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breeds)
}
