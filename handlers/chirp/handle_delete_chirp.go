package chirp

import (
	"errors"
	chirpDomain "github.com/matevskial/chirpyx/domain/chirp"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strconv"
)

func (chirpHandler *ChirpHandler) handleDeleteChirp(w http.ResponseWriter, req *http.Request) {
	authenticationPrincipal, err := handlerutils.GetAuthenticationPrincipalFromRequest(req)
	if err != nil {
		handlerutils.RespondWithUnauthorized(w)
		return
	}

	idStr := req.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "invalid chirp id")
		return
	}

	err = chirpHandler.chirpRepository.DeleteByIdAndAuthorId(id, authenticationPrincipal.UserId)

	if errors.Is(err, chirpDomain.ErrChirpNotFound) {
		handlerutils.RespondWithStatusCode(w, http.StatusForbidden)
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	handlerutils.RespondWithStatusCode(w, http.StatusNoContent)
}
