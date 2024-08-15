package polkaauth

import (
	"errors"
	"net/http"
)

var (
	ErrPolkaNotAuthenticated = errors.New("polka not authenticated")
)

type PolkaAuthenticationPrincipal struct{}

type PolkaAuthenticationService interface {
	Authenticate(req *http.Request) (*PolkaAuthenticationPrincipal, error)
}
