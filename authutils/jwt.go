package authutils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matevskial/chirpyx/configuration"
	"time"
)

const defaultExpiresAfter = 24 * time.Hour

type JwtService struct {
	config *configuration.Configuration
}

func NewJwtService(config *configuration.Configuration) *JwtService {
	return &JwtService{config: config}
}

type JwtGenerateRequest struct {
	UserId           int
	ExpiresInSeconds int
}

func (jwtService *JwtService) GenerateJwtFor(generateRequest JwtGenerateRequest) (string, error) {
	now := time.Now().UTC()
	expiresAfter := getExpiresAfter(generateRequest.ExpiresInSeconds)
	expiresAt := now.Add(expiresAfter)
	claims := jwt.RegisteredClaims{
		Issuer:    jwtService.config.JwtIssuer,
		IssuedAt:  &jwt.NumericDate{Time: now},
		ExpiresAt: &jwt.NumericDate{Time: expiresAt},
		Subject:   fmt.Sprintf("%d", generateRequest.UserId),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtService.config.JwtSecret))
}

func getExpiresAfter(seconds int) time.Duration {
	if seconds == 0 {
		return defaultExpiresAfter
	}

	h := time.Duration(seconds) * time.Second

	if h > defaultExpiresAfter {
		return defaultExpiresAfter
	}

	return h
}
