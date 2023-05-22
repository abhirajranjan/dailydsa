package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// test data
func historyHandler(ctx *gin.Context, db databasebridge) {
	sessionID_string, ok := ctx.Get("sessionID")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	sessionID := sessionID_string.(int)
	history, err := db.GetHistoryBySessionID(sessionID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, history)
}
