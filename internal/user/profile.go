package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type profileResponse struct {
	Name string
}

// handle user profile
func profileHandler(ctx *gin.Context) {
	// test response
	res := profileResponse{
		Name: "abhiraj ranjan",
	}
	ctx.JSON(http.StatusOK, res)
}
