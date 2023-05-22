package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handle user profile
func profileHandler(ctx *gin.Context, db databasebridge) {
	var sessionID int

	sessionID_string, ok := ctx.Get("sessionID")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	sessionID = sessionID_string.(int)
	profile, err := db.GetProfileBySessionID(sessionID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, profile)
}
