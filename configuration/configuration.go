package configuration

import (
	"errors"
	"flag"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

const defaultIssuer = "chirpy"

type Configuration struct {
	IsDevMode   bool
	JwtSecret   string
	JwtIssuer   string
	PolkaApiKey string
}

func Parse() (*Configuration, error) {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		return nil, dotEnvErr
	}

	config := Configuration{}

	isDevModeValueFromFlag := flag.Bool("dev", false, "Enable dev mode")
	flag.Parse()
	config.IsDevMode = *isDevModeValueFromFlag

	isDevModeValueFromEnv, isDevParseBoolErr := strconv.ParseBool(os.Getenv("DEV_MODE"))
	if isDevParseBoolErr != nil {
		isDevModeValueFromEnv = false
	}

	config.IsDevMode = *isDevModeValueFromFlag || isDevModeValueFromEnv

	jwtSecret, jwtSecretErr := parseAndValidateJwtSecret()
	if jwtSecretErr != nil {
		return nil, jwtSecretErr
	}
	config.JwtSecret = jwtSecret

	jwtIssuerFromEnv := os.Getenv("JWT_ISSUER")
	if len(strings.TrimSpace(jwtIssuerFromEnv)) == 0 {
		config.JwtIssuer = defaultIssuer
	} else {
		config.JwtIssuer = jwtIssuerFromEnv
	}

	polkaApiKeyFromEnv := os.Getenv("POLKA_API_KEY")
	config.PolkaApiKey = polkaApiKeyFromEnv

	return &config, nil
}

func parseAndValidateJwtSecret() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if len(strings.TrimSpace(jwtSecret)) == 0 {
		return "", errors.New("invalid jwt secret configuration")
	}
	return jwtSecret, nil
}
