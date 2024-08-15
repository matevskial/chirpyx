package authutils

import (
	"context"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"net/http"
)

type authenticationPrincipalContextKey = string

const authenticationPrincipalContextValue = authenticationPrincipalContextKey("token")

func GetAuthenticationPrincipalFromRequest(req *http.Request) (*authDomain.AuthenticationPrincipal, error) {
	ctx := req.Context()
	authenticationPrincipal, ok := ctx.Value(authenticationPrincipalContextValue).(*authDomain.AuthenticationPrincipal)
	if !ok {
		return nil, authDomain.ErrNotAuthenticated
	}
	return authenticationPrincipal, nil
}

func NewAuthenticatedRequest(req *http.Request, authenticationPrincipal *authDomain.AuthenticationPrincipal) *http.Request {
	oldContext := req.Context()
	newContext := context.WithValue(oldContext, authenticationPrincipalContextValue, authenticationPrincipal)
	return req.WithContext(newContext)
}

func GetBearerTokenString(req *http.Request) (string, error) {
	return getAuthorizationString(req, "Bearer")
}
