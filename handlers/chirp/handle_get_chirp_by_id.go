package chirp

import (
	"errors"
	chirpDomain "github.com/matevskial/chirpyx/domain/chirp"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strconv"
)

func (chirpHandler *ChirpHandler) getChirpById(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "invalid chirp id")
		return
	}

	chirp, err := chirpHandler.chirpRepository.FindById(id)

	if errors.Is(err, chirpDomain.ErrChirpNotFound) {
		handlerutils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	chirpDto := chirpDto{Id: chirp.Id, Body: chirp.Body}
	handlerutils.RespondWithJson(w, http.StatusOK, chirpDto)
}
