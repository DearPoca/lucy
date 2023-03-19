package api

import (
	"net/http"

	"lucy/pkg/log"

	"lucy/pkg/respond"
	"lucy/service/user_service"
	"lucy/utils"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	isExist, err := user_service.CheckAuth(username, password)
	if isExist && err == nil {
		hour := 30 * 24
		token, err := utils.GenerateToken(username, password, hour)
		if err != nil {
			log.Info("User GenerateToken failed", "username", username, "err", err.Error())
			c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUsernameOrPasswordError))
		} else {
			c.SetCookie("token", token, hour*60*60,
				"", "", false, true)
			c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
		}
	} else {
		log.Info("User CheckAuth failed, no such user or password error", "username", username)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUsernameOrPasswordError))
	}
}
