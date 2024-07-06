package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/internal/database"
	"net/http"
	"strconv"
)

type Breed struct {
	ID                       int    `json:"id"`
	Species                  string `json:"species"`
	PetSize                  string `json:"pet_size"`
	Name                     string `json:"name"`
	AverageMaleAdultWeight   int    `json:"average_male_adult_weight"`
	AverageFemaleAdultWeight int    `json:"average_female_adult_weight"`
}

// Get a single breed by ID from the database
func GetBreedHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var breed Breed
	row := db.QueryRow(database.GetBreedByIdQuery, id)
	err := row.Scan(&breed.ID, &breed.Species, &breed.PetSize, &breed.Name, &breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Breed not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(breed)
}

func GetBreedsHanlder(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(database.GetAllBreedsQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var breeds []Breed
	for rows.Next() {
		var b Breed
		err := rows.Scan(&b.ID, &b.Species, &b.PetSize, &b.Name, &b.AverageMaleAdultWeight, &b.AverageFemaleAdultWeight)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		breeds = append(breeds, b)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breeds)
}

func CreateBreedHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var breed Breed
	err := json.NewDecoder(r.Body).Decode(&breed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := database.CreateBreedQuery
	result, err := db.Exec(query, breed.Species, breed.PetSize, breed.Name, breed.AverageMaleAdultWeight, breed.AverageFemaleAdultWeight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	breed.ID = int(lastID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(breed)
}

func UpdateBreedHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var breed Breed
	err := json.NewDecoder(r.Body).Decode(&breed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := database.UpdateBreedQuery
	_, err = db.Exec(query, breed.Species, breed.PetSize, breed.Name, breed.AverageMaleAdultWeight, breed.AverageFemaleAdultWeight, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteBreed(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	query := database.DeleteBreedQuery
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SearchBreedsHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	species := queryParams.Get("species")
	weight := queryParams.Get("weight")
	weightQuery := ""
	weightArgs := []any{}

	if weight != "" {
		weightQuery = " WHERE (average_male_adult_weight = ? OR average_female_adult_weight = ?)"
		weightInt, err := strconv.Atoi(weight)
		if err != nil {
			http.Error(w, "Invalid weight parameter", http.StatusBadRequest)
			return
		}
		weightArgs = append(weightArgs, weightInt, weightInt)
	}

	if species != "" {
		if weightQuery != "" {
			weightQuery += " AND species=?"
		} else {
			weightQuery += " WHERE species=?"
		}
		weightArgs = append(weightArgs, species)
	}

	query := fmt.Sprintf("SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds %s", weightQuery)
	rows, err := db.Query(query, weightArgs...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var breeds []Breed
	for rows.Next() {
		var b Breed
		err := rows.Scan(&b.ID, &b.Species, &b.PetSize, &b.Name, &b.AverageMaleAdultWeight, &b.AverageFemaleAdultWeight)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		breeds = append(breeds, b)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breeds)
}
