package auth

import (
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"net/http"
)

type AuthenticationHandler struct {
	Path                  string
	userRepository        userDomain.UserRepository
	authenticationService authDomain.AuthenticationService
	refreshTokenService   authDomain.RefreshTokenService
}

func NewAuthenticationHandler(
	path string,
	userRepository userDomain.UserRepository,
	authenticationService authDomain.AuthenticationService,
	refreshTokenService authDomain.RefreshTokenService,
) *AuthenticationHandler {
	return &AuthenticationHandler{
		Path:                  path,
		userRepository:        userRepository,
		authenticationService: authenticationService,
		refreshTokenService:   refreshTokenService,
	}
}

func (authenticationHandler *AuthenticationHandler) LoginHandler() http.Handler {
	return http.HandlerFunc(authenticationHandler.handleUserLogin)
}

func (authenticationHandler *AuthenticationHandler) RefreshTokenHandler() http.Handler {
	return http.HandlerFunc(authenticationHandler.handleRefreshToken)
}

func (authenticationHandler *AuthenticationHandler) RevokeRefreshTokenHandler() http.Handler {
	return http.HandlerFunc(authenticationHandler.handleRevokeRefreshToken)
}
