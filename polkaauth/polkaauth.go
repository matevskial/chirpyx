package polkaauth

import (
	"github.com/matevskial/chirpyx/authutils"
	"github.com/matevskial/chirpyx/configuration"
	polkaauthDomain "github.com/matevskial/chirpyx/domain/polkaauth"
	"net/http"
)

type defaultPolkaAuthenticationService struct {
	config *configuration.Configuration
}

func (p *defaultPolkaAuthenticationService) Authenticate(req *http.Request) (*polkaauthDomain.PolkaAuthenticationPrincipal, error) {
	apiKey, err := authutils.GetApiKeyString(req)
	if err != nil {
		return nil, polkaauthDomain.ErrPolkaNotAuthenticated
	}

	if apiKey != p.config.PolkaApiKey {
		return nil, polkaauthDomain.ErrPolkaNotAuthenticated
	}

	return &polkaauthDomain.PolkaAuthenticationPrincipal{}, nil
}

func NewPolkaAuthenticationService(config *configuration.Configuration) polkaauthDomain.PolkaAuthenticationService {
	return &defaultPolkaAuthenticationService{config: config}
}
