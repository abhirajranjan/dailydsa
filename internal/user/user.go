package user

import (
	"github.com/abhirajranjan/dailydsa/internal/auth"
	"github.com/abhirajranjan/dailydsa/internal/permissions"
	"github.com/gin-gonic/gin"
)

// add user routes
//
// param: gin router group inside which user uri lives
func Register(group *gin.RouterGroup) {
	group.GET("/profile", auth.ValidateAuthHandler(permissions.User), profileHandler)
	group.GET("/history", auth.ValidateAuthHandler(permissions.User), historyHandler)
	group.POST("/create", createUserHandler)
}
