package v1

import (
	"net/http"

	"lucy/pkg/respond"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	u := struct {
		Name string
	}{Name: "todo"}
	c.JSON(http.StatusOK, respond.ResSuccess(u))
}
