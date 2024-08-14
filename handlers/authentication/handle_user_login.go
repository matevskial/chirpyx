package authentication

import (
	"encoding/json"
	"errors"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type userLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (authenticationHandler *AuthenticationHandler) handleUserLogin(w http.ResponseWriter, req *http.Request) {
	userLoginRequest := userLoginRequest{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userLoginRequest)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
		return
	}

	user, err := authenticationHandler.userRepository.GetUserWithPasswordByEmail(userLoginRequest.Email)
	if errors.Is(err, userDomain.ErrUserNotFound) {
		handlerutils.RespondWithError(w, http.StatusUnauthorized, "User not found or password mismatch")
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	passwordMatchError := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(userLoginRequest.Password))
	if passwordMatchError != nil {
		handlerutils.RespondWithError(w, http.StatusUnauthorized, "User not found or password mismatch")
		return
	}

	loggedUser := loggedUserResponseDto{
		Id:    user.Id,
		Email: user.Email,
	}
	handlerutils.RespondWithJson(w, http.StatusOK, loggedUser)
}
