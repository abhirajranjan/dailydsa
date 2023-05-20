package user

import (
	"net/http"

	"github.com/abhirajranjan/dailydsa/internal/database"
	"github.com/gin-gonic/gin"
)

// handle user profile
func profileHandler(ctx *gin.Context) {
	var sessionID int

	sessionID_string, ok := ctx.Get("sessionID")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	sessionID = sessionID_string.(int)
	profile, err := database.GetProfileBySessionID(sessionID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, profile)
}
