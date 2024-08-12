package chirp

import "errors"

var (
	ErrChirpNotFound = errors.New("chirp not found")
)

type Chirp struct {
	Id   int
	Body string
}

type ChirpRepository interface {
	Create(body string) (Chirp, error)
	FindAll() ([]Chirp, error)
	FindById(id int) (Chirp, error)
}
