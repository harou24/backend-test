package database

const (
	GetBreedByIdQuery = "SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE id=?"
	GetAllBreedsQuery = "SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds"
)
