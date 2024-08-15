package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matevskial/chirpyx/configuration"
	"github.com/matevskial/chirpyx/domain/auth"
	"time"
)

type JwtService struct {
	config *configuration.Configuration
}

func NewJwtService(config *configuration.Configuration) *JwtService {
	return &JwtService{config: config}
}

type JwtGenerateRequest struct {
	UserId int
}

func (jwtService *JwtService) GenerateJwtFor(generateRequest JwtGenerateRequest) (string, error) {
	now := time.Now().UTC()
	expiresAfter := getExpiresAfter()
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

func (jwtService *JwtService) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtService.config.JwtSecret), nil
	})
}

func getExpiresAfter() time.Duration {
	return auth.AccessTokenExpiresAfterInHours * time.Hour
}
