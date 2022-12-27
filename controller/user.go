package controller

import (
	"log"
	"net/http"

	"lucy/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		log.Printf("An error occurred while registering: %s", err.Error())
		c.JSON(http.StatusOK, ResUnknownError)
		return
	}

	if models.IsUserExisted(body.Username) {
		log.Printf("User %s register, but existed", body.Username)
		c.JSON(http.StatusOK, ResUserExisted)
		return
	}

	models.CreateUser(body.Username, body.Password)
	c.JSON(http.StatusOK, ResSuccess)
}
