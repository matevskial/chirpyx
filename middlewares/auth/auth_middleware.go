package auth

import (
	"github.com/matevskial/chirpyx/authutils"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type AuthenticationMiddleware struct {
	authenticationService authDomain.AuthenticationService
}

func NewAuthenticationMiddleware(authenticationService authDomain.AuthenticationService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{authenticationService: authenticationService}
}

func (am *AuthenticationMiddleware) AuthenticatedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authenticationPrincipal, err := am.authenticationService.Authenticate(req)
		if err != nil {
			handlerutils.RespondWithUnauthorized(w)
			return
		}

		newReq := authutils.NewAuthenticatedRequest(req, authenticationPrincipal)
		next.ServeHTTP(w, newReq)
	})
}
