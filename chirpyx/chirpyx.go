package chirpyx

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type validateChirpRequest struct {
	Body string `json:"body"`
}

type validateChirpErrorResponse struct {
	Error string `json:"error"`
}

type validateChirpResponse struct {
	Valid bool `json:"valid"`
}

func Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("POST /validate_chirp", func(w http.ResponseWriter, req *http.Request) {
		validateChirpRequest, err := decodeValidateChirpRequest(req.Body)
		if err != nil {
			log.Printf("Error decoding body: %s", err)
			errorResponseDto := validateChirpErrorResponse{Error: "Something went wrong"}
			response, err := encodeReponse(&errorResponseDto)
			if err != nil {
				log.Printf("Error marshalling JSON: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}

		switch isValidChirp(validateChirpRequest) {
		case true:
			responseDto := validateChirpResponse{Valid: true}
			response, err := encodeReponse(&responseDto)
			if err != nil {
				log.Printf("Error marshalling JSON: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		case false:
			errorResponseDto := validateChirpErrorResponse{Error: "Chirp is too long"}
			response, err := encodeReponse(&errorResponseDto)
			if err != nil {
				log.Printf("Error marshalling JSON: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
		}
	})

	return httpServeMux
}

func isValidChirp(request *validateChirpRequest) bool {
	return len(request.Body) <= 140
}

func encodeReponse[T any](responseDto *T) ([]byte, error) {
	response, err := json.Marshal(responseDto)
	return response, err
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
