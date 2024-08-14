package user

import (
	"encoding/json"
	"errors"
	"github.com/matevskial/chirpyx/auth"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strconv"
)

func (userHandler *UserHandler) handleUpdateUser(w http.ResponseWriter, req *http.Request) {
	token, ok := auth.GetTokenFromContext(req.Context())
	if !ok {
		handlerutils.RespondWithUnauthorized(w)
		return
	}

	userIdStr, err := token.Claims.GetSubject()
	if err != nil {
		handlerutils.RespondWithUnauthorized(w)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
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

	otherUserWithSameEmailExists, err := userHandler.userRepository.ExistsByEmailAndIdIsNot(userUpdateRequest.Email, userId)
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

	user, err := userHandler.userRepository.Update(userId, userUpdateRequest.Email, hashedPassword)
	if errors.Is(err, userDomain.ErrUserNotFound) {
		handlerutils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	userUpdateResponse := userCreateUpdateResponse{Id: user.Id, Email: user.Email}
	handlerutils.RespondWithJson(w, http.StatusOK, userUpdateResponse)
}
