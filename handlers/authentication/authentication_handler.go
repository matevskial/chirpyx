package authentication

import (
	"github.com/matevskial/chirpyx/auth"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"net/http"
)

type loggedUserResponseDto struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthenticationHandler struct {
	Path           string
	userRepository userDomain.UserRepository
	jwtService     *auth.JwtService
}

func NewAuthenticationHandler(path string, userRepository userDomain.UserRepository, jwtService *auth.JwtService) *AuthenticationHandler {
	return &AuthenticationHandler{Path: path, userRepository: userRepository, jwtService: jwtService}
}

func (authenticationHandler *AuthenticationHandler) LoginHandler() http.Handler {
	return http.HandlerFunc(authenticationHandler.handleUserLogin)
}
