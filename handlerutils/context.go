package handlerutils

import (
	"context"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"github.com/matevskial/chirpyx/polkaauth"
	"net/http"
)

type tokenContextKey = string

const tokenContextKeyValue = tokenContextKey("token")

type polkaAuthenticationPrincipalContextKey = string

const polkaAuthenticationPrincipalContextKeyValue = polkaAuthenticationPrincipalContextKey("polkaAuthenticationPrincipal")

func GetAuthenticationPrincipalFromRequest(req *http.Request) (*authDomain.AuthenticationPrincipal, error) {
	ctx := req.Context()
	authenticationPrincipal, ok := ctx.Value(tokenContextKeyValue).(*authDomain.AuthenticationPrincipal)
	if !ok {
		return nil, authDomain.ErrNotAuthenticated
	}
	return authenticationPrincipal, nil
}

func NewAuthenticatedRequest(req *http.Request, authenticationPrincipal *authDomain.AuthenticationPrincipal) *http.Request {
	oldContext := req.Context()
	newContext := context.WithValue(oldContext, tokenContextKeyValue, authenticationPrincipal)
	return req.WithContext(newContext)
}

func NewPolkaAuthenticatedRequest(req *http.Request, polkaAuthenticationPrincipal *polkaauth.PolkaAuthenticationPrincipal) *http.Request {
	oldContext := req.Context()
	newContext := context.WithValue(oldContext, polkaAuthenticationPrincipalContextKeyValue, polkaAuthenticationPrincipal)
	return req.WithContext(newContext)
}
