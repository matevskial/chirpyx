package polka

import (
	"encoding/json"
	"errors"
	userDomain "github.com/matevskial/chirpyx/domain/user"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type polkaEvent string

const (
	upgraded = "user.upgraded"
)

type webhookRequest struct {
	Event polkaEvent `json:"event"`
	Data  userData   `json:"data"`
}

type userData struct {
	UserId int `json:"user_id"`
}

func (p *PolkaHandler) handleWebhooks(w http.ResponseWriter, req *http.Request) {
	webhookRequest := webhookRequest{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&webhookRequest)
	if err != nil {
		handlerutils.RespondWithError(w, http.StatusBadRequest, "Couldn't decode body")
		return
	}

	if webhookRequest.Event != upgraded {
		handlerutils.RespondWithStatusCode(w, http.StatusNoContent)
		return
	}

	user, err := p.userRepository.FindById(webhookRequest.Data.UserId)
	if errors.Is(err, userDomain.ErrUserNotFound) {
		handlerutils.RespondWithStatusCode(w, http.StatusNotFound)
		return
	}

	if user.IsChirpyRed {
		handlerutils.RespondWithStatusCode(w, http.StatusNoContent)
		return
	}

	err = p.userRepository.UpgradeToChirpyRed(user.Id)
	if errors.Is(err, userDomain.ErrUserNotFound) {
		handlerutils.RespondWithStatusCode(w, http.StatusNotFound)
		return
	} else if err != nil {
		handlerutils.RespondWithInternalServerError(w)
	}

	handlerutils.RespondWithStatusCode(w, http.StatusNoContent)
}
