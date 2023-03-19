package api

import (
	"net/http"

	"lucy/pkg/log"

	"lucy/pkg/respond"
	"lucy/service/user_service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	telephone := c.Query("telephone")

	if user_service.IsUserExisted(username) {
		log.Info("User register, but existed", "username", username)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUserExisted))
		return
	}

	err := user_service.CreateUser(username, password, email, telephone)
	if err != nil {
		log.Info("User register failed, err: %s", "username", username, "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeParamInvalid))
		return
	}
	log.Info("User register failed", "username", username, "email", email, "telephone", telephone)
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
}
