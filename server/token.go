package server

import "github.com/dgrijalva/jwt-go"

type token struct {
	ID    uint
	Email string
	*jwt.StandardClaims
}
