package authentication

import (
	"github.com/matevskial/chirpyx/authutils"
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
	jwtService     *authutils.JwtService
}

func NewAuthenticationHandler(path string, userRepository userDomain.UserRepository, jwtService *authutils.JwtService) *AuthenticationHandler {
	return &AuthenticationHandler{Path: path, userRepository: userRepository, jwtService: jwtService}
}

func (authenticationHandler *AuthenticationHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST"+" "+authenticationHandler.Path, authenticationHandler.handleUserLogin)
	return httpServeMux
}
