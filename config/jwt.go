package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("fgfdgdgdfgrgr98b")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
