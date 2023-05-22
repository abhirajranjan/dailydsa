package server

import (
	"github.com/abhirajranjan/dailydsa/internal/admin"
	"github.com/abhirajranjan/dailydsa/internal/daily"
	"github.com/abhirajranjan/dailydsa/internal/model"
	"github.com/abhirajranjan/dailydsa/internal/permissions"
	"github.com/abhirajranjan/dailydsa/internal/user"
	"github.com/gin-gonic/gin"
)

type databasebridge interface {
	CreateUser(jwt *model.JWT) (int, error)
	GetUserRolesBySessionID(sessionID int) (permissions.Permissions, error)
	GetProfileBySessionID(sessionID int) (*model.UserProfileModel, error)
	GetHistoryBySessionID(int) (*model.HistoryResponse, error)
}

func Serve(db databasebridge) *gin.Engine {
	engine := gin.New()

	// group uri
	userUri := engine.Group("/user")
	adminUri := engine.Group("/admin")
	dailyUri := engine.Group("/daily")

	// register endpoints
	user.Register(userUri, db)
	admin.Register(adminUri, db)
	daily.Register(dailyUri, db)
	return engine

}
