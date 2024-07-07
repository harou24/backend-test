package domain

type BreedRepository interface {
	GetBreedByID(int) (*Breed, error)
	GetAllBreeds() ([]Breed, error)
	CreateBreed(Breed) (*Breed, error)
	UpdateBreed(int, Breed) error
	DeleteBreed(int) error
	SearchBreeds(criteria map[string]any) ([]Breed, error)
}
