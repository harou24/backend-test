package persistence

import (
	"database/sql"
	"fmt"
	"github.com/japhy-tech/backend-test/internal/domain"
)

type MysqlBreedRepository struct {
	db *sql.DB
}

func NewMysqlBreedRepository(db *sql.DB) domain.BreedRepository {
	return &MysqlBreedRepository{db: db}
}

func (r *MysqlBreedRepository) GetBreedByID(id int) (*domain.Breed, error) {
	query := "SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var breed domain.Breed
	err := row.Scan(&breed.ID, &breed.Species, &breed.PetSize, &breed.Name, &breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrBreedNotFound
		}
		return nil, err
	}

	return &breed, nil
}
func (r *MysqlBreedRepository) GetAllBreeds(page, limit int) ([]domain.Breed, int, error) {
	offset := (page - 1) * limit

	// Query total count of items
	var totalCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM breeds").Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Query paginated data
	rows, err := r.db.Query("SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var breeds []domain.Breed
	for rows.Next() {
		var breed domain.Breed
		err := rows.Scan(&breed.ID, &breed.Species, &breed.PetSize, &breed.Name, &breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight)
		if err != nil {
			return nil, 0, err
		}
		breeds = append(breeds, breed)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return breeds, totalCount, nil
}

func (r *MysqlBreedRepository) CreateBreed(breed domain.Breed) (*domain.Breed, error) {
	result, err := r.db.Exec(CreateBreedQuery, breed.Species, breed.PetSize, breed.Name, breed.AverageMaleAdultWeight, breed.AverageFemaleAdultWeight)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	breed.ID = int(lastID)

	return &breed, nil
}

func (r *MysqlBreedRepository) UpdateBreed(id int, breed domain.Breed) error {
	_, err := r.db.Exec(UpdateBreedQuery, breed.Species, breed.PetSize, breed.Name, breed.AverageMaleAdultWeight, breed.AverageFemaleAdultWeight, id)
	return err
}

func (r *MysqlBreedRepository) DeleteBreed(id int) error {
	_, err := r.db.Exec(DeleteBreedQuery, id)
	return err
}

func (r *MysqlBreedRepository) SearchBreeds(criteria map[string]any) ([]domain.Breed, error) {
	query := "SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE 1=1"
	args := []interface{}{}

	for key, value := range criteria {
		switch key {
		case "species":
			query += " AND species = ?"
			args = append(args, value)
		case "weight":
			weight, ok := value.(int)
			if !ok {
				return nil, fmt.Errorf("invalid weight type: %T", value)
			}
			query += " AND (average_male_adult_weight = ? OR average_female_adult_weight = ?)"
			args = append(args, weight, weight)
		default:
			return nil, fmt.Errorf("unknown search criteria: %s", key)
		}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var breeds []domain.Breed
	for rows.Next() {
		var breed domain.Breed
		err := rows.Scan(&breed.ID, &breed.Species, &breed.PetSize, &breed.Name, &breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, breed)
	}

	return breeds, nil
}
