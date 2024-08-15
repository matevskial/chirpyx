package user

import (
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/middlewares/auth"
	"net/http"
)

type UserHandler struct {
	Path                     string
	userRepository           userDomain.UserRepository
	authenticationMiddleware *auth.AuthenticationMiddleware
}

type userCreateUpdateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userCreateUpdateResponse struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	IsChirpyRed bool   `json:"is_chirpy_red"`
}

func NewUserHandler(path string, userRepository userDomain.UserRepository, authenticationMiddleware *auth.AuthenticationMiddleware) *UserHandler {
	return &UserHandler{Path: path, userRepository: userRepository, authenticationMiddleware: authenticationMiddleware}
}

func (userHandler *UserHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST"+" "+userHandler.Path, userHandler.handleCreateUser)
	httpServeMux.Handle("PUT"+" "+userHandler.Path, userHandler.authenticationMiddleware.AuthenticatedHandler(http.HandlerFunc(userHandler.handleUpdateUser)))
	return httpServeMux
}
