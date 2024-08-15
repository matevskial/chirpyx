package user

import (
	"encoding/json"
	"errors"
	"github.com/matevskial/chirpyx/auth"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

func (userHandler *UserHandler) handleUpdateUser(w http.ResponseWriter, req *http.Request) {
	authenticationPrincipal, err := handlerutils.GetAuthenticationPrincipalFromRequest(req)
	if err != nil {
		handlerutils.RespondWithUnauthorized(w)
		return
	}

	userUpdateRequest := userCreateUpdateRequest{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&userUpdateRequest)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
		return
	}

	otherUserWithSameEmailExists, err := userHandler.userRepository.ExistsByEmailAndIdIsNot(userUpdateRequest.Email, authenticationPrincipal.UserId)
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}
	if otherUserWithSameEmailExists {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "User with provided email already exists")
		return
	}

	hashedPassword, err := auth.HashPassword(userUpdateRequest.Password)
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	user, err := userHandler.userRepository.Update(authenticationPrincipal.UserId, userUpdateRequest.Email, hashedPassword)
	if errors.Is(err, userDomain.ErrUserNotFound) {
		handlerutils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	userUpdateResponse := userCreateUpdateResponse{Id: user.Id, Email: user.Email, IsChirpyRed: user.IsChirpyRed}
	handlerutils.RespondWithJson(w, http.StatusOK, userUpdateResponse)
}
