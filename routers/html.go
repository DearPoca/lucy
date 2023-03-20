package routers

import (
	"net/http"

	"lucy/middleware/jwt"
	"lucy/pkg/respond"
	"lucy/service/media_service"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"username": username,
	})
}

func userinfo(c *gin.Context) {
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	c.HTML(http.StatusOK, "userinfo.tmpl", gin.H{
		"username": username,
	})
}

func login(c *gin.Context) {
	c.SetCookie("", "", -1, "", "", false, true)
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{})
}

func playWebrtc(c *gin.Context) {
	liveId := c.Query("live_id")
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	live, err := media_service.GetLiveById(liveId)
	if err != nil {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeLiveNotFound))
		return
	}
	c.HTML(http.StatusOK, "play_webrtc.tmpl", gin.H{
		"username":   username,
		"webrtc_url": live.WebrtcUrl,
	})
}

func playFlv(c *gin.Context) {
	liveId := c.Query("live_id")
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	live, err := media_service.GetLiveById(liveId)
	if err != nil {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeLiveNotFound))
		return
	}
	c.HTML(http.StatusOK, "play_flv.tmpl", gin.H{
		"username": username,
		"flv_url":  live.FlvUrl,
	})
	return
}

func newLive(c *gin.Context) {
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}

	c.HTML(http.StatusOK, "new_live.tmpl", gin.H{
		"username": username,
	})
}
