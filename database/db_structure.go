package database

import (
	chirpDomain "github.com/matevskial/chirpyx/domain/chirp"
	userDomain "github.com/matevskial/chirpyx/domain/user"
)

type DBStructure struct {
	Chirps               map[int]chirpDomain.Chirp `json:"chirps"`
	Users                map[int]User              `json:"users"`
	RefreshTokensByUsers map[int]RefreshToken      `json:"refresh_tokens"`
	ChirpIdSeq           int                       `json:"idSeq"`
	UserIdSeq            int                       `json:"userIdSeq"`
}

func newDbStructure() DBStructure {
	return DBStructure{
		Chirps:               make(map[int]chirpDomain.Chirp),
		Users:                make(map[int]User),
		RefreshTokensByUsers: make(map[int]RefreshToken),
		ChirpIdSeq:           1,
		UserIdSeq:            1,
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

func (s *DBStructure) updateUser(id int, email string, hashedPassword string) (userDomain.User, error) {
	user, exists := s.Users[id]
	if !exists {
		return userDomain.User{}, userDomain.ErrUserNotFound
	}
	user.Email = email
	user.HashedPassword = hashedPassword
	s.Users[id] = user
	return userDomain.User{Id: user.Id, Email: user.Email}, nil
}

func (s *DBStructure) deleteChirpByIdAndAuthorId(chirpId int, authorId int) error {
	chirp, exists := s.Chirps[chirpId]
	if !exists || chirp.AuthorId != authorId {
		return chirpDomain.ErrChirpNotFound
	}
	delete(s.Chirps, chirpId)
	return nil
}

func (s *DBStructure) getChirpById(id int) (chirpDomain.Chirp, error) {
	chirp, exists := s.Chirps[id]
	if !exists {
		return chirpDomain.Chirp{}, chirpDomain.ErrChirpNotFound
	}
	return chirp, nil
}
