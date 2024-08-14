package database

import (
	"encoding/json"
	"errors"
	chirpDomain "github.com/matevskial/chirpyx/domain/chirp"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"os"
	"sync"
)

type JsonFileDB struct {
	path string
	mux  *sync.RWMutex
}

type User struct {
	Id             int
	Email          string
	HashedPassword []byte
}

type DBStructure struct {
	Chirps     map[int]chirpDomain.Chirp `json:"chirps"`
	Users      map[int]User              `json:"users"`
	ChirpIdSeq int                       `json:"idSeq"`
	UserIdSeq  int                       `json:"userIdSeq"`
}

func newDbStructure() DBStructure {
	return DBStructure{
		Chirps:     make(map[int]chirpDomain.Chirp),
		Users:      make(map[int]User),
		ChirpIdSeq: 1,
		UserIdSeq:  1,
	}
}

func (s *DBStructure) addChirp(chirp chirpDomain.Chirp) {
	s.Chirps[chirp.Id] = chirp
	s.ChirpIdSeq++
}

func (s *DBStructure) addUser(user User) userDomain.User {
	s.Users[user.Id] = user
	s.UserIdSeq++
	return userDomain.User{Id: user.Id, Email: user.Email}
}

func NewDB(path string, shouldTruncateExistingFile bool) (*JsonFileDB, error) {
	db := &JsonFileDB{path: path, mux: &sync.RWMutex{}}
	err := db.ensureDB(shouldTruncateExistingFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *JsonFileDB) CreateChirp(body string) (chirpDomain.Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return chirpDomain.Chirp{}, err
	}
	chirp := chirpDomain.Chirp{Id: dbStructure.ChirpIdSeq, Body: body}
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

func (db *JsonFileDB) CreateUser(email string, hashedPassword []byte) (userDomain.User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return userDomain.User{}, err
	}
	userToAdd := User{Id: dbStructure.UserIdSeq, Email: email, HashedPassword: hashedPassword}
	addedUser := dbStructure.addUser(userToAdd)
	err = db.writeDB(dbStructure)
	if err != nil {
		return userDomain.User{}, err
	}
	return addedUser, nil
}

func (db *JsonFileDB) ExistsUserByEmail(email string) (bool, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return false, err
	}

	for _, value := range dbStructure.Users {
		if value.Email == email {
			return true, nil
		}
	}

	return false, nil
}

func (db *JsonFileDB) GetUserByEmail(email string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, value := range dbStructure.Users {
		if value.Email == email {
			return value, nil
		}
	}

	return User{}, userDomain.ErrUserNotFound
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
