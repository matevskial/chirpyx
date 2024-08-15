package auth

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/matevskial/chirpyx/database"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"time"
)

type refreshTokenJsonFileDatabaseService struct {
	db *database.JsonFileDB
}

func (rs *refreshTokenJsonFileDatabaseService) CreateRefreshToken(userId int) (string, error) {
	refreshTokenString, err := generateRefreshToken()
	if err != nil {
		return "", err
	}
	err = rs.db.SetRefreshToken(userId, refreshTokenString, getExpiresAt())
	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}

func generateRefreshToken() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}

func getExpiresAt() time.Time {
	return time.Now().UTC().AddDate(0, 0, authDomain.RefreshTokenExpiresAfterInDays)
}

func (rs *refreshTokenJsonFileDatabaseService) RevokeRefreshToken(refreshTokenString string) error {
	return rs.db.RevokeRefreshToken(refreshTokenString)
}

func (rs *refreshTokenJsonFileDatabaseService) GetRefreshToken(refreshTokenString string) (authDomain.RefreshToken, error) {
	refreshToken, err := rs.db.GetRefreshToken(refreshTokenString)
	if err != nil {
		return authDomain.RefreshToken{}, err
	}
	if refreshToken.IsExpired() {
		return authDomain.RefreshToken{}, authDomain.ErrRefreshTokenExpired
	}
	return refreshToken, nil
}

func NewRefreshTokenService(db *database.JsonFileDB) authDomain.RefreshTokenService {
	return &refreshTokenJsonFileDatabaseService{db: db}
}
