package polka

import (
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type PolkaHandler struct {
	userRepository userDomain.UserRepository
}

func NewPolkaHandler(userRepository userDomain.UserRepository) *PolkaHandler {
	return &PolkaHandler{userRepository: userRepository}
}

func (p *PolkaHandler) Handler(pathPrefix string) http.Handler {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc(handlerutils.PostRequestPath(pathPrefix, "webhooks"), p.handleWebhooks)
	return httpServeMux
}
