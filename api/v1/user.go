package v1

import (
	"net/http"

	"lucy/pkg/log"

	"lucy/middleware/jwt"
	"lucy/pkg/respond"
	"lucy/service/user_service"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	username, ok := c.Get(jwt.KeyOfUsername)
	if !ok {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	u, err := user_service.GetUserInfo(username.(string))

	if err != nil {
		log.Info("Get user info failed", "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeGetUserInfoFailed))
	} else {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, u))
	}
}
