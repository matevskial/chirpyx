package chirp

import (
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

func (chirpHandler *ChirpHandler) handleGetChirps(w http.ResponseWriter, _ *http.Request) {
	chirps, err := chirpHandler.chirpRepository.FindAll()
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}
	chirpDtos := make([]chirpDto, len(chirps))
	for i, value := range chirps {
		chirpDtos[i] = chirpDto{Id: value.Id, Body: value.Body}
	}
	handlerutils.RespondWithJson(w, http.StatusOK, chirpDtos)
}
