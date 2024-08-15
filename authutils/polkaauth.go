package authutils

import (
	"context"
	polkaauthDomain "github.com/matevskial/chirpyx/domain/polkaauth"
	"net/http"
)

type polkaAuthenticationPrincipalContextKey = string

const polkaAuthenticationPrincipalContextKeyValue = polkaAuthenticationPrincipalContextKey("polkaAuthenticationPrincipal")

func NewPolkaAuthenticatedRequest(req *http.Request, polkaAuthenticationPrincipal *polkaauthDomain.PolkaAuthenticationPrincipal) *http.Request {
	oldContext := req.Context()
	newContext := context.WithValue(oldContext, polkaAuthenticationPrincipalContextKeyValue, polkaAuthenticationPrincipal)
	return req.WithContext(newContext)
}

func GetApiKeyString(req *http.Request) (string, error) {
	return getAuthorizationString(req, "ApiKey")
}
