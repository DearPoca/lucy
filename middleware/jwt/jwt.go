package jwt

import (
	"net/http"
	"time"

	"lucy/pkg/log"

	"lucy/utils"

	"github.com/gin-gonic/gin"
)

const HeaderAuthorizationKey = "Authorization"
const KeyOfUsername = "username"

func JWT(c *gin.Context) {
	redirect := func() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.Abort()
	}
	token, err := c.Cookie("token")
	if token == "" || err != nil {
		log.Info("no cookie, redirect", "token", token, "err", err.Error())
		redirect()
		return
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Info("token parse failed", "err", err.Error())
		redirect()
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		redirect()
		return
	}
	c.Set(KeyOfUsername, claims.Username)
	c.Next()
}
