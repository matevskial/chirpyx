package chirp

import (
	"encoding/json"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strings"
)

type chirpCreateRequest struct {
	Body string `json:"body"`
}

type chirpCreateResponse = chirpDto

func (chirpHandler *ChirpHandler) handleCreateChirp(w http.ResponseWriter, req *http.Request) {
	chirpCreateRequest := chirpCreateRequest{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&chirpCreateRequest)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
		return
	}

	switch isValidChirp(chirpCreateRequest) {
	case true:
		cleanedBody := cleanBody(chirpCreateRequest)
		createdChirp, err := chirpHandler.chirpRepository.Create(cleanedBody)
		if err != nil {
			handlerutils.RespondWithInternalServerError(w)
			return
		}
		responseDto := chirpCreateResponse{Id: createdChirp.Id, Body: createdChirp.Body}
		handlerutils.RespondWithJson(w, http.StatusCreated, responseDto)
	case false:
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
	}
}

func isValidChirp(request chirpCreateRequest) bool {
	return len(request.Body) <= 140
}

func cleanBody(chirpRequest chirpCreateRequest) string {
	parts := strings.Split(chirpRequest.Body, " ")
	for i := 0; i < len(parts); i++ {
		if isProfaneWord(parts[i]) {
			parts[i] = "****"
		}
	}
	return strings.Join(parts, " ")
}

func isProfaneWord(s string) bool {
	lowerCaseStr := strings.ToLower(s)
	return lowerCaseStr == "kerfuffle" || lowerCaseStr == "sharbert" || lowerCaseStr == "fornax"
}
