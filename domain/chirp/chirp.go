package chirp

import (
	"errors"
	"github.com/matevskial/chirpyx/common"
)

var (
	ErrChirpNotFound = errors.New("chirp not found")
)

type Chirp struct {
	Id       int
	Body     string
	AuthorId int
}

type ChirpFiltering struct {
	AuthorId int
}

type ChirpRepository interface {
	Create(body string, authorId int) (Chirp, error)
	FindBy(filtering ChirpFiltering, sorting common.Sorting) ([]Chirp, error)
	FindById(id int) (Chirp, error)
	DeleteByIdAndAuthorId(id int, authorId int) error
}
