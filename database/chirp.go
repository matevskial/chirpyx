package database

import chirpDomain "github.com/matevskial/chirpyx/domain/chirp"

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

func (db *JsonFileDB) GetChirps() ([]chirpDomain.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	chirps := make([]chirpDomain.Chirp, len(dbStructure.Chirps))
	i := 0
	for _, value := range dbStructure.Chirps {
		chirps[i] = value
		i++
	}
	return chirps, nil
}
