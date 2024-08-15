package chirp

import (
	"github.com/matevskial/chirpyx/domain/chirp"
	authMiddleware "github.com/matevskial/chirpyx/middlewares/auth"
	"net/http"
)

type ChirpHandler struct {
	chirpRepository          chirp.ChirpRepository
	authenticationMiddleware *authMiddleware.AuthenticationMiddleware
}

func NewChirpHandler(chirpRepository chirp.ChirpRepository, authenticationMiddleware *authMiddleware.AuthenticationMiddleware) *ChirpHandler {
	return &ChirpHandler{chirpRepository: chirpRepository, authenticationMiddleware: authenticationMiddleware}
}

func (chirpHandler *ChirpHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.Handle("POST /chirps", chirpHandler.authenticationMiddleware.AuthenticatedHandler(http.HandlerFunc(chirpHandler.handleCreateChirp)))
	httpServeMux.HandleFunc("GET /chirps", chirpHandler.handleGetChirps)
	httpServeMux.HandleFunc("GET /chirps/{id}", chirpHandler.getChirpById)
	httpServeMux.Handle("DELETE /chirps/{id}", chirpHandler.authenticationMiddleware.AuthenticatedHandler(http.HandlerFunc(chirpHandler.handleDeleteChirp)))
	return httpServeMux
}

type chirpDto struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	AuthorId int    `json:"author_id"`
}
