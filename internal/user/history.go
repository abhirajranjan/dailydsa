package user

import (
	"net/http"

	"github.com/abhirajranjan/dailydsa/internal/database"
	"github.com/gin-gonic/gin"
)

// test data
func historyHandler(ctx *gin.Context) {
	sessionID_string, ok := ctx.Get("sessionID")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	sessionID := sessionID_string.(int)
	history, err := database.GetHistoryBySessionID(sessionID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, history)
}
