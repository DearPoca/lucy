package v1

import (
	"net/http"

	"lucy/middleware/jwt"
	"lucy/pkg/respond"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	tmp, ok := c.Get(jwt.KeyOfUsername)
	if !ok {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	username := tmp.(string)

	u := struct {
		Name string `json:"name"`
	}{Name: username}

	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, u))
}
