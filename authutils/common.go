package authutils

import (
	"errors"
	"net/http"
	"strings"
)

func getAuthorizationString(req *http.Request, prefix string) (string, error) {
	tokenHeader := req.Header.Get("Authorization")
	finalPrefix := prefix + " "
	tokenString := strings.TrimPrefix(tokenHeader, finalPrefix)
	if isInvalidTokenHeaderValue(tokenString, tokenHeader) {
		return "", errors.New("invalid authorization header")
	}

	if len(strings.TrimSpace(tokenString)) == 0 {
		return "", errors.New("invalid authorization header")
	}

	return tokenString, nil
}

func isInvalidTokenHeaderValue(tokenString, tokenHeader string) bool {
	return len(tokenString) == len(tokenHeader)
}
