package admin

import (
	"github.com/gin-gonic/gin"
)

type databasebridge interface{}

func Register(group *gin.RouterGroup, db databasebridge) {

}
