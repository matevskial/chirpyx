package user

import (
	"encoding/json"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type userCreateRequest struct {
	Email string `json:"email"`
}

type userCreateResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func (userHandler *UserHandler) handleCreateUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	userCreateRequest := userCreateRequest{}
	err := decoder.Decode(&userCreateRequest)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
		return
	}

	user, err := userHandler.userRepository.Create(userCreateRequest.Email)

	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	userCreateResponse := userCreateResponse{Id: user.Id, Email: user.Email}
	handlerutils.RespondWithJson(w, http.StatusCreated, userCreateResponse)
}
