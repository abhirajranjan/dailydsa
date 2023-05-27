package user

import (
	"github.com/abhirajranjan/dailydsa/internal/auth"
	"github.com/abhirajranjan/dailydsa/internal/model"
	"github.com/abhirajranjan/dailydsa/internal/permissions"
	"github.com/gin-gonic/gin"
)

type databasebridge interface {
	CreateUser(jwt *model.JWT) (int, error)
	GetUserRolesBySessionID(sessionID int) (permissions.Permissions, error)
	GetProfileBySessionID(sessionID int) (*model.ModelUser, error)
	GetHistoryBySessionID(int) ([]model.ModelSubmission, error)
}

// add user routes
//
// param: gin router group inside which user uri lives
func Register(group *gin.RouterGroup, db databasebridge) {
	group.GET("/profile", auth.ValidateAuthHandler(permissions.User), func(ctx *gin.Context) {
		profileHandler(ctx, db)
	})
	group.GET("/history", auth.ValidateAuthHandler(permissions.User), func(ctx *gin.Context) {
		historyHandler(ctx, db)
	})
	group.POST("/create", func(ctx *gin.Context) {
		createUserHandler(ctx, db)
	})
}
