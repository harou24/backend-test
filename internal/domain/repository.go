package domain

type BreedRepository interface {
	GetBreedByID(int) (*Breed, error)
	GetAllBreeds(page int, limit int) ([]Breed, int, error)
	CreateBreed(Breed) (*Breed, error)
	UpdateBreed(int, Breed) error
	DeleteBreed(int) error
	SearchBreeds(criteria map[string]any) ([]Breed, error)
}
