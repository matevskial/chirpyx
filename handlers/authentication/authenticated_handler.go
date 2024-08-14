package authentication

import (
	"errors"
	"github.com/matevskial/chirpyx/authutils"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	jwtService *authutils.JwtService
}

func NewAuthenticationMiddleware(jwtService *authutils.JwtService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{jwtService: jwtService}
}

func (am *AuthenticationMiddleware) AuthenticatedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tokenString, err := getTokenString(req)
		if err != nil {
			handlerutils.RespondWithUnauthorized(w)
			return
		}

		token, err := am.jwtService.ParseToken(tokenString)
		if err != nil {
			handlerutils.RespondWithUnauthorized(w)
			return
		}

		oldContext := req.Context()
		newContext := authutils.NewContextWithTokenValue(oldContext, token)
		next.ServeHTTP(w, req.WithContext(newContext))
	})
}

func getTokenString(req *http.Request) (string, error) {
	tokenHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
	if isInvalidTokenHeaderValue(tokenString, tokenHeader) {
		return "", errors.New("invalid token header")
	}

	if len(strings.TrimSpace(tokenString)) == 0 {
		return "", errors.New("invalid token header")
	}

	return tokenString, nil
}

func isInvalidTokenHeaderValue(tokenString, tokenHeader string) bool {
	return len(tokenString) == len(tokenHeader)
}
