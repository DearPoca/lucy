package api

import (
	"log"
	"net/http"

	"lucy/models"
	"lucy/pkg/respond"
	"lucy/utils"

	"github.com/beego/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
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
	ok, _ := valid.Valid(&body)

	if ok {
		isExist, err := models.CheckAuth(body.Username, body.Password)
		if isExist && err == nil {
			token, err := utils.GenerateToken(body.Username, body.Password)
			if err != nil {
				log.Printf("user [%s] GenerateToken failed: %s", body.Username, err.Error())
				c.JSON(http.StatusOK, respond.ResUsernameOrPasswordError())
			} else {
				res := respond.ResSuccess()
				data := make(map[string]interface{})
				data["token"] = token
				res.Data = data
				c.JSON(http.StatusOK, res)
			}
		} else {
			log.Printf("user [%s] CheckAuth failed, no such user", body.Username)
			c.JSON(http.StatusOK, respond.ResUsernameOrPasswordError())
		}
	} else {
		c.JSON(http.StatusOK, respond.ResUsernameOrPasswordError())
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
}
