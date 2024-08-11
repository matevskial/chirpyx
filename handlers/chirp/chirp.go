package chirp

import (
	"encoding/json"
	"github.com/matevskial/chirpyx/domain/chirp"
	"github.com/matevskial/chirpyx/handlerutils"
	"io"
	"net/http"
	"strings"
)

type chirpCreateRequest struct {
	Body string `json:"body"`
}

type chirpDto struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type chirpCreateResponse = chirpDto

func cleanBody(chirpRequest *chirpCreateRequest) string {
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

func Handler(chirpRepository chirp.ChirpRepository) http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST /chirps", func(w http.ResponseWriter, req *http.Request) {
		chirpCreateRequest, err := decodeChirpCreateRequest(req.Body)
		if err != nil {
			handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
			return
		}

		switch isValidChirp(chirpCreateRequest) {
		case true:
			cleanedBody := cleanBody(chirpCreateRequest)
			createdChirp, err := chirpRepository.Create(cleanedBody)
			if err != nil {
				handlerutils.RespondWithInternalServerError(w)
				return
			}
			responseDto := chirpCreateResponse{Id: createdChirp.Id, Body: createdChirp.Body}
			handlerutils.RespondWithJson(w, http.StatusCreated, responseDto)
		case false:
			handlerutils.RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
		}
	})

	httpServeMux.HandleFunc("GET /chirps", func(w http.ResponseWriter, req *http.Request) {
		chirps, err := chirpRepository.FindAll()
		if err != nil {
			handlerutils.RespondWithInternalServerError(w)
			return
		}
		chirpDtos := make([]chirpDto, len(chirps))
		for i, value := range chirps {
			chirpDtos[i] = chirpDto{Id: value.Id, Body: value.Body}
		}
		handlerutils.RespondWithJson(w, http.StatusOK, chirpDtos)
	})

	return httpServeMux
}

func isValidChirp(request *chirpCreateRequest) bool {
	return len(request.Body) <= 140
}

func decodeChirpCreateRequest(body io.ReadCloser) (*chirpCreateRequest, error) {
	decoder := json.NewDecoder(body)
	chirpCreateRequest := chirpCreateRequest{}
	err := decoder.Decode(&chirpCreateRequest)
	if err != nil {
		return nil, err
	}
	return &chirpCreateRequest, nil
}
