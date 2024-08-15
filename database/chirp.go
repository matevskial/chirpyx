package database

import (
	"github.com/matevskial/chirpyx/common"
	chirpDomain "github.com/matevskial/chirpyx/domain/chirp"
	"sort"
)

func (db *JsonFileDB) CreateChirp(body string, authorId int) (chirpDomain.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return chirpDomain.Chirp{}, err
	}
	chirp := chirpDomain.Chirp{Id: dbStructure.ChirpIdSeq, Body: body, AuthorId: authorId}
	dbStructure.addChirp(chirp)
	err = db.writeDB(dbStructure)
	if err != nil {
		return chirpDomain.Chirp{}, err
	}
	return chirp, nil
}

func (db *JsonFileDB) GetChirps(filtering chirpDomain.ChirpFiltering, sorting common.Sorting) ([]chirpDomain.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	chirps := make([]chirpDomain.Chirp, 0)
	for _, value := range dbStructure.Chirps {
		if chirpMatchesFiltering(value, filtering) {
			chirps = append(chirps, value)
		}
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sorting.Direction == common.Asc {
			return chirps[i].Id < chirps[j].Id
		}
		return chirps[j].Id < chirps[i].Id
	})

	return chirps, nil
}

func chirpMatchesFiltering(chirp chirpDomain.Chirp, filtering chirpDomain.ChirpFiltering) bool {
	return filtering.AuthorId == 0 || chirp.AuthorId == filtering.AuthorId
}

func (db *JsonFileDB) GetChirpById(id int) (chirpDomain.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return chirpDomain.Chirp{}, err
	}

	return dbStructure.getChirpById(id)
}

func (db *JsonFileDB) DeleteChirpByIdAndAuthorId(id int, authorId int) error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	err = dbStructure.deleteChirpByIdAndAuthorId(id, authorId)
	if err != nil {
		return err
	}

	return db.writeDB(dbStructure)
}
