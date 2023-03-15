package jwt

import (
	"net/http"
	"strings"
	"time"

	"lucy/utils"

	"github.com/gin-gonic/gin"
)

const HeaderAuthorizationKey = "Authorization"
const KeyOfUsername = "username"

func JWT(c *gin.Context) {
	token := c.GetHeader(HeaderAuthorizationKey)
	redirect := func() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.Abort()
	}
	if token == "" || !strings.HasPrefix(token, "Bearer") {
		redirect()
		return
	}

	// Remove "Bearer "
	token = token[7:]

	claims, err := utils.ParseToken(token)
	if err != nil {
		redirect()
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		redirect()
		return
	}
	c.Set(KeyOfUsername, claims.Username)
	c.Next()
}
