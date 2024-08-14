package authentication

import (
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"net/http"
)

type loggedUserResponseDto struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type AuthenticationHandler struct {
	Path           string
	userRepository userDomain.UserRepository
}

func NewAuthenticationHandler(path string, userRepository userDomain.UserRepository) *AuthenticationHandler {
	return &AuthenticationHandler{Path: path, userRepository: userRepository}
}

func (authenticationHandler *AuthenticationHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST"+" "+authenticationHandler.Path, authenticationHandler.handleUserLogin)
	return httpServeMux
}
