package chirpyx

import (
	"encoding/json"
	"github.com/matevskial/chirpyx/handlerutils"
	"io"
	"net/http"
	"strings"
)

type validateChirpRequest struct {
	Body string `json:"body"`
}

type validateChirpResponse struct {
	Valid       bool   `json:"valid"`
	CleanedBody string `json:"cleaned_body"`
}

func cleanBody(chirpRequest *validateChirpRequest) string {
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

func Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST /validate_chirp", func(w http.ResponseWriter, req *http.Request) {
		validateChirpRequest, err := decodeValidateChirpRequest(req.Body)
		if err != nil {
			handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
			return
		}

		switch isValidChirp(validateChirpRequest) {
		case true:
			cleanedBody := cleanBody(validateChirpRequest)
			responseDto := validateChirpResponse{Valid: true, CleanedBody: cleanedBody}
			handlerutils.RespondWithJson(w, http.StatusOK, responseDto)
		case false:
			handlerutils.RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
		}
	})

	return httpServeMux
}

func isValidChirp(request *validateChirpRequest) bool {
	return len(request.Body) <= 140
}

func decodeValidateChirpRequest(body io.ReadCloser) (*validateChirpRequest, error) {
	decoder := json.NewDecoder(body)
	validateChirpRequest := validateChirpRequest{}
	err := decoder.Decode(&validateChirpRequest)
	if err != nil {
		return nil, err
	}
	return &validateChirpRequest, nil
}
