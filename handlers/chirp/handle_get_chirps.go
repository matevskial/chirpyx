package chirp

import (
	"github.com/matevskial/chirpyx/domain/chirp"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strconv"
)

func (chirpHandler *ChirpHandler) handleGetChirps(w http.ResponseWriter, req *http.Request) {
	chirpFiltering := chirp.ChirpFiltering{}
	err := setChirpFiltering(req, &chirpFiltering)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "error parsing query parameters")
	}

	chirps, err := chirpHandler.chirpRepository.FindBy(chirpFiltering)
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

func setChirpFiltering(req *http.Request, c *chirp.ChirpFiltering) error {
	authorIdStr := req.URL.Query().Get("author_id")
	if authorIdStr == "" {
		c.AuthorId = 0
		return nil
	}
	authorId, err := strconv.Atoi(authorIdStr)
	if err != nil {
		return err
	}
	c.AuthorId = authorId
	return nil
}
