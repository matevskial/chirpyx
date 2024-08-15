package authentication

import (
	"errors"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

func (authenticationHandler *AuthenticationHandler) handleRevokeRefreshToken(w http.ResponseWriter, req *http.Request) {
	refreshTokenString, err := handlerutils.GetBearerTokenString(req)
	if err != nil {
		handlerutils.RespondWithUnauthorized(w)
		return
	}

	err = authenticationHandler.refreshTokenService.RevokeRefreshToken(refreshTokenString)
	if errors.Is(err, authDomain.ErrRefreshTokenNotFound) {
		handlerutils.RespondWithStatusCode(w, http.StatusNoContent)
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
		return
	}

	handlerutils.RespondWithStatusCode(w, http.StatusNoContent)
}
