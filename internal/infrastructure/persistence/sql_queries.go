package persistence

const (
	GetBreedByIdQuery = "SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE id=?"
	GetAllBreedsQuery = "SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM breeds"
	CreateBreedQuery  = "INSERT INTO breeds (species, pet_size, name, average_male_adult_weight, average_female_adult_weight) VALUES (?, ?, ?, ?, ?)"
	UpdateBreedQuery  = "UPDATE breeds SET species=?, pet_size=?, name=?, average_male_adult_weight=?, average_female_adult_weight=? WHERE id=?"
	DeleteBreedQuery  = "DELETE FROM breeds WHERE id=?"
)