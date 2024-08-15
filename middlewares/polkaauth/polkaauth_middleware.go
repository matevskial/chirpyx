package polkaauth

import (
	"github.com/matevskial/chirpyx/authutils"
	polkaauthDomain "github.com/matevskial/chirpyx/domain/polkaauth"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type PolkaAuthenticationMiddleware struct {
	polkaAuthenticationService polkaauthDomain.PolkaAuthenticationService
}

func NewPolkaAuthenticationMiddleware(polkaAuthenticationService polkaauthDomain.PolkaAuthenticationService) *PolkaAuthenticationMiddleware {
	return &PolkaAuthenticationMiddleware{polkaAuthenticationService: polkaAuthenticationService}
}

func (am *PolkaAuthenticationMiddleware) AuthenticatedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authenticationPrincipal, err := am.polkaAuthenticationService.Authenticate(req)
		if err != nil {
			handlerutils.RespondWithUnauthorized(w)
			return
		}

		newReq := authutils.NewPolkaAuthenticatedRequest(req, authenticationPrincipal)
		next.ServeHTTP(w, newReq)
	})
}
