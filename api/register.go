package api

import (
	"log"
	"net/http"

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
		log.Printf("User [%s] register, but existed", username)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUserExisted))
		return
	}

	err := user_service.CreateUser(username, password, email, telephone)
	if err != nil {
		log.Printf("User [%s] register failed, err: %s", username, err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeParamInvalid))
		return
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
}
