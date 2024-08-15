package auth

import (
	"errors"
	"time"
)

var (
	ErrRefreshTokenExpired  = errors.New("refresh token expired")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)

const (
	RefreshTokenExpiresAfterInDays = 60
	AccessTokenExpiresAfterInHours = 1
)

type RefreshToken struct {
	RefreshTokenString string
	UserId             int
	ExpiresAt          time.Time
}

func (rt *RefreshToken) IsExpired() bool {
	now := time.Now().UTC()
	return now.After(rt.ExpiresAt)
}

type RefreshTokenService interface {
	CreateRefreshToken(userId int) (string, error)
	RevokeRefreshToken(refreshTokenString string) error
	GetRefreshToken(refreshTokenString string) (RefreshToken, error)
}
