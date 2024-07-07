package domain

import "fmt"

// ErrBreedNotFound is returned when a breed is not found in the database.
var ErrBreedNotFound = fmt.Errorf("breed not found")
