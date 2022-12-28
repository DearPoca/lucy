package jwt

import (
	"net/http"
	"strings"
	"time"

	"lucy/pkg/respond"
	"lucy/utils"

	"github.com/gin-gonic/gin"
)

const HeaderAuthorizationKey = "Authorization"
const KeyOfUsername = "username"

func JWT(c *gin.Context) {
	token := c.GetHeader(HeaderAuthorizationKey)
	if token == "" || !strings.HasPrefix(token, "Bearer") {
		c.JSON(http.StatusUnauthorized, respond.CreateRespond(respond.CodeParamInvalid))
		c.Abort()
		return
	}

	// Remove "Bearer "
	token = token[7:]

	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, respond.CreateRespond(respond.CodeAuthCheckTokenFail))
		c.Abort()
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		c.JSON(http.StatusUnauthorized, respond.CreateRespond(respond.CodeAuthTimeout))
		c.Abort()
		return
	}
	c.Set(KeyOfUsername, claims.Username)
	c.Next()
}
