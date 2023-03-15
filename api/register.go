package api

import (
	"log"
	"net/http"

	"lucy/pkg/respond"
	"lucy/service/user_service"

	"github.com/beego/beego/validation"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		Telephone string `json:"telephone"`
	}
	if err := c.BindJSON(&body); err != nil {
		log.Printf("An error occurred while registering: %s", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(&body)

	if !ok || err != nil {
		log.Printf("An error occurred while registering: %s", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}

	if user_service.IsUserExisted(body.Username) {
		log.Printf("User %s register, but existed", body.Username)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUserExisted))
		return
	}

	user_service.CreateUser(body.Username, body.Password)
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
}
