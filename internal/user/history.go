package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// per element format
type questionDetails struct {
	Link       string    `json:"link"`
	Difficulty string    `json:"difficulty"`
	Issolved   bool      `json:"issolved"`
	Solution   string    `json:"solution"`
	DateTime   time.Time `json:"datetime"`
}

// response format
type historyResponse struct {
	Count    int             `json:"count"`
	Elements questionDetails `json:"child"`
}

// test data
func historyHandler(ctx *gin.Context) {
	res := historyResponse{
		Count: 1,
		Elements: questionDetails{
			Link:       "testlink/test",
			Difficulty: "intermediate",
			Issolved:   false,
			Solution:   "",
			DateTime:   time.Now(),
		},
	}
	ctx.JSON(http.StatusOK, res)
}
