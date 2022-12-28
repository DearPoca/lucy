package api

import (
	"log"
	"net/http"

	"lucy/models"
	"lucy/pkg/respond"

	"github.com/beego/beego/validation"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		log.Printf("An error occurred while registering: %s", err.Error())
		c.JSON(http.StatusOK, respond.ResUnknownError())
		return
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(&body)

	if !ok || err != nil {
		log.Printf("An error occurred while registering: %s", err.Error())
		c.JSON(http.StatusOK, respond.ResUnknownError())
		return
	}

	if models.IsUserExisted(body.Username) {
		log.Printf("User %s register, but existed", body.Username)
		c.JSON(http.StatusOK, respond.ResUserExisted())
		return
	}

	models.CreateUser(body.Username, body.Password)
	c.JSON(http.StatusOK, respond.ResSuccess())
}
