package auth

import (
	"flag"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// signing key used to verify jwt
	key string

	// signing used to parse jwt
	signing_alg = []string{"HS256", "HS384", "HS512"}
)

func init() {
	flag.StringVar(&key, "hashing key", "", "signing key used to verify jwt")
}

func ParseJwt(jwttoken string) map[string]interface{} {
	validAlg := jwt.WithValidMethods(signing_alg)

	token, err := jwt.Parse(jwttoken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(key), nil
	}, validAlg)

	if err != nil {
		// TODO: log jwt parse error
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		// TODO: log token claims type error
		return nil
	}
}
