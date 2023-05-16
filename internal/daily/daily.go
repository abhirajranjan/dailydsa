package daily

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionFormat struct {
	Link       string `json:"link"`
	Name       string `json:"name"`
	Difficulty string `json:"difficulty"`
}

func Register(group *gin.RouterGroup) {
	group.GET("", GetCurrentQuestion)
}

func GetCurrentQuestion(ctx *gin.Context) {
	//TODO: fetch current question
	res := QuestionFormat{
		Name:       "2sum",
		Link:       "https://leetcode.com/problems/two-sum/",
		Difficulty: "easy",
	}
	ctx.JSON(http.StatusOK, res)
}
