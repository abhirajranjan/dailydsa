package server

import (
	"github.com/abhirajranjan/dailydsa/internal/admin"
	"github.com/abhirajranjan/dailydsa/internal/daily"
	"github.com/abhirajranjan/dailydsa/internal/user"
	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	engine := gin.New()

	// group uri
	userUri := engine.Group("/user")
	adminUri := engine.Group("/admin")
	dailyUri := engine.Group("/daily")

	// register endpoints
	user.Register(userUri)
	admin.Register(adminUri)
	daily.Register(dailyUri)
	return engine

}
