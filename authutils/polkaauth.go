package authutils

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKeyString(req *http.Request) (string, error) {
	tokenHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(tokenHeader, "ApiKey ")
	if isInvalidTokenHeaderValue(tokenString, tokenHeader) {
		return "", errors.New("invalid token header")
	}

	if len(strings.TrimSpace(tokenString)) == 0 {
		return "", errors.New("invalid token header")
	}

	return tokenString, nil
}

func isInvalidTokenHeaderValue(tokenString, tokenHeader string) bool {
	return len(tokenString) == len(tokenHeader)
}
