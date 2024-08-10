package main

import (
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
)

type healthHandler struct{}

func (healthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	responseString := "OK"
	handlerutils.RespondWithText(w, http.StatusOK, responseString)
}
