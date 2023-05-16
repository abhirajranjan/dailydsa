package auth

import (
	"fmt"
	"net/http"

	"github.com/abhirajranjan/dailydsa/internal/helper"
	"github.com/gin-gonic/gin"
)

const (
	// cookieTag is mapped to jwt in cookies of requests
	cookieTag = "auth"
)

func ValidateAuthHandler(reqperm helper.Permissions) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateAuth(ctx, reqperm) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "user login required")
		}
	}
}

// validate user cookie
//
// returns true if user is valid else false
func validateAuth(c *gin.Context, reqperm helper.Permissions) bool {
	var authparam string
	var err error

	authparam, err = c.Cookie(cookieTag)
	// if no cookie found return false
	if err == http.ErrNoCookie {
		return false
	}

	// TODO: add jwt decoder and jwt flow

	// just to use variable
	fmt.Printf("authparam: %v\n", authparam)
	return true
}
