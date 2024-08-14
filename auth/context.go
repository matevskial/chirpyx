package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type tokenContextKey = string

const tokenContextKeyValue = tokenContextKey("token")

func GetTokenFromContext(ctx context.Context) (*jwt.Token, bool) {
	token, ok := ctx.Value(tokenContextKeyValue).(*jwt.Token)
	return token, ok
}

func NewContextWithTokenValue(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, tokenContextKeyValue, token)
}
