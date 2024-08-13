package user

import (
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"net/http"
)

type UserHandler struct {
	Path           string
	userRepository userDomain.UserRepository
}

func NewUserHandler(path string, userRepository userDomain.UserRepository) *UserHandler {
	return &UserHandler{Path: path, userRepository: userRepository}
}

func (userHandler *UserHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST"+" "+userHandler.Path, userHandler.handleCreateUser)
	return httpServeMux
}
