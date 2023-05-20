package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhirajranjan/dailydsa/internal/database"
	"github.com/abhirajranjan/dailydsa/internal/permissions"
	"github.com/gin-gonic/gin"
)

const (
	// cookieTag is mapped to jwt in cookies of requests
	cookieTag = "auth"
)

func ValidateAuthHandler(reqperm permissions.Permissions) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validateAuth(ctx, reqperm) {
			if len(ctx.Errors) == 0 {
				ctx.Status(http.StatusUnauthorized)
				ctx.Abort()
				return
			}

			err := ctx.Errors[len(ctx.Errors)-1]
			status := err.Meta.(int)

			if err.Type == gin.ErrorTypePublic {
				ctx.AbortWithStatusJSON(status, err.Err.Error())
			} else {
				ctx.AbortWithError(status, err.Err)
			}
		}

		ctx.Next()
	}
}

// validate user cookie
//
// returns true if user is valid for required permissions
func validateAuth(ctx *gin.Context, requiredperm permissions.Permissions) bool {
	var sessionID_string string
	var sessionID int
	var err error

	// check if request is from frontend and can frontend access it
	if isFrontend(ctx) && requiredperm.Has(permissions.Frontend) {
		return true
	}

	sessionID_string, err = ctx.Cookie(cookieTag)
	// if no cookie found return false
	if err == http.ErrNoCookie {
		ctx.Error(ginErr(http.StatusUnauthorized, "invalid session ID", gin.ErrorTypePublic))
		return false
	}

	sessionID, err = strconv.Atoi(sessionID_string)
	if err != nil {
		ctx.Error(ginErr(http.StatusUnauthorized, "invalid session ID", gin.ErrorTypePublic))
		return false
	}

	ctx.Set("sessionID", sessionID)

	userroles, err := database.GetUserRolesBySessionID(sessionID)
	if err != nil {
		ctx.Error(ginErr(http.StatusInternalServerError, "database failure", gin.ErrorTypePrivate))
		return false
	}

	if !userroles.Has(requiredperm) {
		ctx.Error(ginErr(http.StatusUnauthorized, "unauthorised", gin.ErrorTypePublic))
		return false
	}

	return true
}

// perform logic checks if frontend
func isFrontend(c *gin.Context) bool {
	if a := c.GetHeader("auth"); a != "" {
		return true
	}
	return false
}

func ginErr(status int, errstring string, Type gin.ErrorType) *gin.Error {
	return &gin.Error{Err: fmt.Errorf(errstring), Type: Type, Meta: status}
}
