package user

import (
	"net/http"

	"github.com/abhirajranjan/dailydsa/internal/auth"
	"github.com/abhirajranjan/dailydsa/internal/database"
	"github.com/abhirajranjan/dailydsa/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type jwtrequest struct {
	JwtToken string `json:"jwt"`
}

func createUserHandler(ctx *gin.Context) {
	var Jwtrequest jwtrequest
	var jwtres model.JWT

	if err := ctx.Bind(&Jwtrequest); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	token := auth.ParseJwt(Jwtrequest.JwtToken)
	if token == nil {
		ctx.JSON(http.StatusBadRequest, "failed parsing jwt")
		return
	}

	if err := mapToStruct(&jwtres, token); err != nil {
		ctx.JSON(http.StatusBadRequest, "exceeding/limiting tags specified")
	}

	tempid, err := database.CreateUser(&jwtres)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	ctx.JSON(http.StatusAccepted, tempid)
}

func mapToStruct(result *model.JWT, inputMap map[string]interface{}) error {
	config := &mapstructure.DecoderConfig{
		// ErrorUnused:          true,
		// ErrorUnset:           true,
		TagName:              "jwt",
		IgnoreUntaggedFields: true,
		Result:               result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(inputMap)
	if err != nil {
		return err
	}

	return nil
}
