package handlerutils

import (
	"context"
	authDomain "github.com/matevskial/chirpyx/domain/auth"
	"net/http"
)

type tokenContextKey = string

const tokenContextKeyValue = tokenContextKey("token")

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
