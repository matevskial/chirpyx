package database

import (
	"encoding/json"
	"errors"
	"github.com/matevskial/chirpyx/domain/chirp"
	"os"
	"sync"
)

type JsonFileDB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]chirp.Chirp `json:"chirps"`
	IdSeq  int                 `json:"idSeq"`
}

func newDbStructure() DBStructure {
	return DBStructure{Chirps: make(map[int]chirp.Chirp), IdSeq: 1}
}

func (s *DBStructure) addChirp(chrp chirp.Chirp) {
	s.Chirps[chrp.Id] = chrp
	s.IdSeq++
}

func NewDB(path string) (*JsonFileDB, error) {
	db := &JsonFileDB{path: path, mux: &sync.RWMutex{}}
	err := db.ensureDB()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *JsonFileDB) CreateChirp(body string) (chirp.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return chirp.Chirp{}, err
	}
	chrp := chirp.Chirp{Id: dbStructure.IdSeq, Body: body}
	dbStructure.addChirp(chrp)
	err = db.writeDB(dbStructure)
	if err != nil {
		return chirp.Chirp{}, err
	}
	return chrp, nil
}

func (db *JsonFileDB) GetChirps() ([]chirp.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	chirps := make([]chirp.Chirp, len(dbStructure.Chirps))
	i := 0
	for _, value := range dbStructure.Chirps {
		chirps[i] = value
		i++
	}
	return chirps, nil
}

func (db *JsonFileDB) ensureDB() error {
	if _, err := os.Stat(db.path); errors.Is(err, os.ErrNotExist) {
		data, jsonErr := json.Marshal(newDbStructure())
		if jsonErr != nil {
			return jsonErr
		}
		return os.WriteFile(db.path, data, 0666)
	}
	return nil
}

func (db *JsonFileDB) loadDB() (DBStructure, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()
	data, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, err
	}

	dbStructure := DBStructure{}
	err = json.Unmarshal(data, &dbStructure)
	if err != nil {
		return DBStructure{}, err
	}
	return dbStructure, nil
}

func (db *JsonFileDB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()
	data, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
