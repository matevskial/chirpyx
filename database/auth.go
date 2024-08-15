package database

import (
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"time"
)

type RefreshToken struct {
	UserId             int
	RefreshTokenString string
	ExpiresAt          time.Time
}

func (db *JsonFileDB) GetRefreshToken(refreshTokenString string) (authDomain.RefreshToken, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return authDomain.RefreshToken{}, err
	}

	for _, value := range dbStructure.RefreshTokensByUsers {
		if value.RefreshTokenString == refreshTokenString {
			return authDomain.RefreshToken{UserId: value.UserId, RefreshTokenString: value.RefreshTokenString, ExpiresAt: value.ExpiresAt}, nil
		}
	}

	return authDomain.RefreshToken{}, authDomain.ErrRefreshTokenNotFound
}

func (db *JsonFileDB) RevokeRefreshToken(refreshTokenString string) error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	var refreshToken *RefreshToken = nil
	for _, value := range dbStructure.RefreshTokensByUsers {
		if value.RefreshTokenString == refreshTokenString {
			refreshToken = &value
			break
		}
	}

	if refreshToken == nil {
		return authDomain.ErrRefreshTokenNotFound
	}

	delete(dbStructure.RefreshTokensByUsers, refreshToken.UserId)

	return db.writeDB(dbStructure)
}

func (db *JsonFileDB) SetRefreshToken(userId int, refreshTokenString string, expiresAt time.Time) error {
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}

	dbStructure.RefreshTokensByUsers[userId] = RefreshToken{UserId: userId, RefreshTokenString: refreshTokenString, ExpiresAt: expiresAt}

	return db.writeDB(dbStructure)
}
