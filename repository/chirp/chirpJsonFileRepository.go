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

func (r *ChirpJsonFileRepository) Create(body string) (chirp.Chirp, error) {
	return r.db.CreateChirp(body)
}

func (r *ChirpJsonFileRepository) FindAll() ([]chirp.Chirp, error) {
	return r.db.GetChirps()
}
