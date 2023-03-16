package api

import (
	"log"
	"net/http"

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
		token, err := utils.GenerateToken(username, password)
		if err != nil {
			log.Printf("user [%s] GenerateToken failed: %s", username, err.Error())
			c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUsernameOrPasswordError))
		} else {
			c.SetCookie("token", token, 0,
				"", "", false, true)
			c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
		}
	} else {
		log.Printf("user [%s] CheckAuth failed, no such user or password error", username)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUsernameOrPasswordError))
	}
}
