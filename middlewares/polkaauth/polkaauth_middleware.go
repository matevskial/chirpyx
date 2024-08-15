package polkaauth

import (
	"github.com/matevskial/chirpyx/handlerutils"
	"github.com/matevskial/chirpyx/polkaauth"
	"net/http"
)

type PolkaAuthenticationMiddleware struct {
	polkaAuthenticationService polkaauth.PolkaAuthenticationService
}

func NewPolkaAuthenticationMiddleware(polkaAuthenticationService polkaauth.PolkaAuthenticationService) *PolkaAuthenticationMiddleware {
	return &PolkaAuthenticationMiddleware{polkaAuthenticationService: polkaAuthenticationService}
}

func (am *PolkaAuthenticationMiddleware) AuthenticatedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authenticationPrincipal, err := am.polkaAuthenticationService.Authenticate(req)
		if err != nil {
			handlerutils.RespondWithUnauthorized(w)
			return
		}

		newReq := handlerutils.NewPolkaAuthenticatedRequest(req, authenticationPrincipal)
		next.ServeHTTP(w, newReq)
	})
}
