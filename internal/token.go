package app

import jwt "github.com/dgrijalva/jwt-go"

// Token struct declaration
type Token struct {
	ID    uint
	Email string
	*jwt.StandardClaims
}
