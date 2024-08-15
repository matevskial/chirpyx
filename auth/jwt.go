package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matevskial/chirpyx/configuration"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strconv"
	"time"
)

type AuthenticationJwtService struct {
	config *configuration.Configuration
}

func NewAuthenticationJwtService(config *configuration.Configuration) authDomain.AuthenticationService {
	return &AuthenticationJwtService{config: config}
}

func (jwtService *AuthenticationJwtService) GenerateToken(generateRequest authDomain.GenerateTokenRequest) (string, error) {
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

func (jwtService *AuthenticationJwtService) Authenticate(req *http.Request) (*authDomain.AuthenticationPrincipal, error) {
	tokenString, err := handlerutils.GetBearerTokenString(req)
	if err != nil {
		return nil, authDomain.ErrNotAuthenticated
	}

	token, err := jwtService.parseToken(tokenString)
	if err != nil {
		return nil, authDomain.ErrNotAuthenticated
	}

	userIdStr, err := token.Claims.GetSubject()
	if err != nil {
		return nil, authDomain.ErrNotAuthenticated
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, authDomain.ErrNotAuthenticated
	}

	return &authDomain.AuthenticationPrincipal{UserId: userId}, nil
}

func (jwtService *AuthenticationJwtService) parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtService.config.JwtSecret), nil
	})
}

func getExpiresAfter() time.Duration {
	return authDomain.AccessTokenExpiresAfterInHours * time.Hour
}
