package chirp

import (
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

func (r *ChirpJsonFileRepository) FindBy(filtering chirp.ChirpFiltering) ([]chirp.Chirp, error) {
	return r.db.GetChirps(filtering)
}

func (r *ChirpJsonFileRepository) FindById(id int) (chirp.Chirp, error) {
	return r.db.GetChirpById(id)
}

func (r *ChirpJsonFileRepository) DeleteByIdAndAuthorId(id int, authorId int) error {
	return r.db.DeleteChirpByIdAndAuthorId(id, authorId)
}
