package auth

import (
	"github.com/matevskial/chirpyx/auth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type AuthenticationMiddleware struct {
	jwtService *auth.JwtService
}

func NewAuthenticationMiddleware(jwtService *auth.JwtService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{jwtService: jwtService}
}

func (am *AuthenticationMiddleware) AuthenticatedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tokenString, err := handlerutils.GetBearerTokenString(req)
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
		newContext := auth.NewContextWithTokenValue(oldContext, token)
		next.ServeHTTP(w, req.WithContext(newContext))
	})
}
