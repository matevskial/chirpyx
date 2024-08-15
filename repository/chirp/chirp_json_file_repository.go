package chirp

import (
	"github.com/matevskial/chirpyx/common"
	"github.com/matevskial/chirpyx/database"
	"github.com/matevskial/chirpyx/domain/chirp"
)

type ChirpJsonFileRepository struct {
	db *database.JsonFileDB
}

func NewChirpJsonFileRepository(db *database.JsonFileDB) *ChirpJsonFileRepository {
	return &ChirpJsonFileRepository{db: db}
}

func (r *ChirpJsonFileRepository) Create(body string, authorId int) (chirp.Chirp, error) {
	return r.db.CreateChirp(body, authorId)
}

func (r *ChirpJsonFileRepository) FindBy(filtering chirp.ChirpFiltering, sorting common.Sorting) ([]chirp.Chirp, error) {
	return r.db.GetChirps(filtering, sorting)
}

func (r *ChirpJsonFileRepository) FindById(id int) (chirp.Chirp, error) {
	return r.db.GetChirpById(id)
}

func (r *ChirpJsonFileRepository) DeleteByIdAndAuthorId(id int, authorId int) error {
	return r.db.DeleteChirpByIdAndAuthorId(id, authorId)
}
