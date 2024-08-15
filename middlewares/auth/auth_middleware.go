package auth

import (
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type AuthenticationMiddleware struct {
	authenticationService authDomain.AuthenticationService
}

func NewAuthenticationMiddleware(jwtService authDomain.AuthenticationService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{authenticationService: jwtService}
}

func (am *AuthenticationMiddleware) AuthenticatedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authenticationPrincipal, err := am.authenticationService.Authenticate(req)
		if err != nil {
			handlerutils.RespondWithUnauthorized(w)
			return
		}

		newReq := handlerutils.NewAuthenticatedRequest(req, authenticationPrincipal)
		next.ServeHTTP(w, newReq)
	})
}
