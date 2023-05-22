package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJwt(jwttoken string) map[string]interface{} {
	validAlg := jwt.WithValidMethods(auth.SigningAlg)

	token, err := jwt.Parse(jwttoken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(auth.Key), nil
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
