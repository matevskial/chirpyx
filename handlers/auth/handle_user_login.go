package auth

import (
	"encoding/json"
	"errors"
	"github.com/matevskial/chirpyx/authutils"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type userLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loggedUserResponseDto struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	IsChirpyRed  bool   `json:"is_chirpy_red"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
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

	passwordMatchError := authutils.ComparePasswordWithHash(userLoginRequest.Password, user.HashedPassword)
	if passwordMatchError != nil {
		handlerutils.RespondWithError(w, http.StatusUnauthorized, "User not found or password mismatch")
		return
	}

	token, err := authenticationHandler.authenticationService.GenerateToken(authDomain.GenerateTokenRequest{UserId: user.Id})

	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	refreshTokenString, err := authenticationHandler.refreshTokenService.CreateRefreshToken(user.Id)

	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	loggedUser := loggedUserResponseDto{
		Id:           user.Id,
		Email:        user.Email,
		IsChirpyRed:  user.IsChirpyRed,
		Token:        token,
		RefreshToken: refreshTokenString,
	}
	handlerutils.RespondWithJson(w, http.StatusOK, loggedUser)
}
