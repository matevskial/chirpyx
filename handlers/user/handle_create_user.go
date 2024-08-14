package user

import (
	"encoding/json"
	"github.com/matevskial/chirpyx/authutils"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

func (userHandler *UserHandler) handleCreateUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	userCreateRequest := userCreateUpdateRequest{}
	err := decoder.Decode(&userCreateRequest)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
		return
	}

	userExists, err := userHandler.userRepository.ExistsByEmail(userCreateRequest.Email)
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}
	if userExists {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "User with provided email already exists")
		return
	}

	hashedPassword, err := authutils.HashPassword(userCreateRequest.Password)
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	user, err := userHandler.userRepository.Create(userCreateRequest.Email, hashedPassword)
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	userCreateResponse := userCreateUpdateResponse{Id: user.Id, Email: user.Email}
	handlerutils.RespondWithJson(w, http.StatusCreated, userCreateResponse)
}
