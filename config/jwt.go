package config

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("rahasia")

type JWTclaim struct {
	Id   int
	Name string
	jwt.RegisteredClaims
}
