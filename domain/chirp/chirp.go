package chirp

import "errors"

var (
	ErrChirpNotFound = errors.New("chirp not found")
)

type Chirp struct {
	Id       int
	Body     string
	AuthorId int
}

type ChirpRepository interface {
	Create(body string, authorId int) (Chirp, error)
	FindAll() ([]Chirp, error)
	FindById(id int) (Chirp, error)
}
