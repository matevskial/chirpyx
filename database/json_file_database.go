package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type JsonFileDB struct {
	path string
	mux  *sync.RWMutex
}

func NewDB(path string, shouldTruncateExistingFile bool) (*JsonFileDB, error) {
	db := &JsonFileDB{path: path, mux: &sync.RWMutex{}}
	err := db.ensureDB(shouldTruncateExistingFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *JsonFileDB) ensureDB(shouldTruncateExistingFile bool) error {
	if _, err := os.Stat(db.path); shouldTruncateExistingFile || errors.Is(err, os.ErrNotExist) {
		return db.writeDB(newDbStructure())
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

	return os.WriteFile(db.path, data, 0666)
}
