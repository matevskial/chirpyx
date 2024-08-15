package user

import (
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"github.com/matevskial/chirpyx/middlewares/auth"
	"net/http"
	"path"
)

type UserHandler struct {
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

func NewUserHandler(userRepository userDomain.UserRepository, authenticationMiddleware *auth.AuthenticationMiddleware) *UserHandler {
	return &UserHandler{userRepository: userRepository, authenticationMiddleware: authenticationMiddleware}
}

func (userHandler *UserHandler) Handler(pathPrefix string) http.Handler {
	httpServeMux := http.NewServeMux()
	cleanedPathPrefix := path.Clean(pathPrefix)
	httpServeMux.HandleFunc(handlerutils.PostRequestPath(cleanedPathPrefix), userHandler.handleCreateUser)
	httpServeMux.Handle(handlerutils.PutRequestPath(cleanedPathPrefix), userHandler.authenticationMiddleware.AuthenticatedHandler(http.HandlerFunc(userHandler.handleUpdateUser)))
	return httpServeMux
}
