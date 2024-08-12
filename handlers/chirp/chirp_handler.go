package chirp

import (
	"github.com/matevskial/chirpyx/domain/chirp"
	"net/http"
)

type ChirpHandler struct {
	chirpRepository chirp.ChirpRepository
}

func NewChirpHandler(chirpRepository chirp.ChirpRepository) *ChirpHandler {
	return &ChirpHandler{chirpRepository: chirpRepository}
}

func (chirpHandler *ChirpHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST /chirps", chirpHandler.handleCreateChirp)
	httpServeMux.HandleFunc("GET /chirps", chirpHandler.handleGetChirps)
	httpServeMux.HandleFunc("GET /chirps/{id}", chirpHandler.getChirpById)
	return httpServeMux
}

type chirpDto struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}
