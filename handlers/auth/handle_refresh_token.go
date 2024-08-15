package auth

import (
	"errors"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type refreshTokenResponse struct {
	Token string `json:"token"`
}

func (authenticationHandler *AuthenticationHandler) handleRefreshToken(w http.ResponseWriter, req *http.Request) {
	refreshTokenString, err := handlerutils.GetBearerTokenString(req)
	if err != nil {
		handlerutils.RespondWithUnauthorized(w)
		return
	}

	refreshToken, err := authenticationHandler.refreshTokenService.GetRefreshToken(refreshTokenString)
	if errors.Is(err, authDomain.ErrRefreshTokenNotFound) || errors.Is(err, authDomain.ErrRefreshTokenExpired) {
		handlerutils.RespondWithUnauthorized(w)
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	token, err := authenticationHandler.authenticationService.GenerateToken(authDomain.GenerateTokenRequest{UserId: refreshToken.UserId})
	if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	refreshTokenResponse := refreshTokenResponse{Token: token}
	handlerutils.RespondWithJson(w, http.StatusOK, refreshTokenResponse)
}
