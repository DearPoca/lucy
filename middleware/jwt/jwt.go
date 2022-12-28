package jwt

import (
	"net/http"
	"strings"
	"time"

	"lucy/pkg/respond"
	"lucy/utils"

	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer") {
		c.JSON(http.StatusUnauthorized, respond.ResParamInvalid())
		c.Abort()
		return
	}

	// Remove "Bearer "
	token = token[7:]

	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, respond.ResAuthCheckTokenFail())
		c.Abort()
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		c.JSON(http.StatusUnauthorized, respond.ResAuthTimeout())
		c.Abort()
		return
	}
	c.Next()
}
