package polkaauth

import (
	"errors"
	"github.com/matevskial/chirpyx/authutils"
	"github.com/matevskial/chirpyx/configuration"
	"net/http"
)

var (
	ErrPolkaNotAuthenticated = errors.New("polka not authenticated")
)

type PolkaAuthenticationPrincipal struct{}

type PolkaAuthenticationService interface {
	Authenticate(req *http.Request) (*PolkaAuthenticationPrincipal, error)
}

type defaultPolkaAuthenticationService struct {
	config *configuration.Configuration
}

func (p *defaultPolkaAuthenticationService) Authenticate(req *http.Request) (*PolkaAuthenticationPrincipal, error) {
	apiKey, err := authutils.GetApiKeyString(req)
	if err != nil {
		return nil, ErrPolkaNotAuthenticated
	}

	if apiKey != p.config.PolkaApiKey {
		return nil, ErrPolkaNotAuthenticated
	}

	return &PolkaAuthenticationPrincipal{}, nil
}

func NewPolkaAuthenticationService(config *configuration.Configuration) PolkaAuthenticationService {
	return &defaultPolkaAuthenticationService{config: config}
}
