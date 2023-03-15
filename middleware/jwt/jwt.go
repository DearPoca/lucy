package jwt

import (
	"log"
	"net/http"
	"time"

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
		log.Printf("no cookie, redirect, token: %s, err: %s", token, err.Error())
		redirect()
		return
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Printf("token parse failed, err: %s", err.Error())
		redirect()
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		redirect()
		return
	}
	c.Set(KeyOfUsername, claims.Username)
	c.Next()
}
