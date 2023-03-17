package routers

import (
	"fmt"
	"net/http"

	"lucy/middleware/jwt"
	"lucy/pkg/setting"
	"lucy/srs"
	"lucy/utils"

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
	streams := srs.GetStreams()
	roomId := c.Query("room_id")
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	for i, _ := range streams {
		if streams[i].Id == roomId {
			url := fmt.Sprintf("webrtc://%s%s", setting.SrsSetting.Ip, streams[i].Url)
			c.HTML(http.StatusOK, "play_webrtc.tmpl", gin.H{
				"username":   username,
				"webrtc_url": url,
			})
			return
		}
	}
	c.String(http.StatusOK, "room no found!")
}

func playFlv(c *gin.Context) {
	streams := srs.GetStreams()
	roomId := c.Query("room_id")
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	for i, _ := range streams {
		if streams[i].Id == roomId {
			url := fmt.Sprintf("http://%s:%s%s.flv",
				setting.SrsSetting.Ip,
				setting.SrsSetting.NginxHttpPort,
				streams[i].Url)
			c.HTML(http.StatusOK, "play_flv.tmpl", gin.H{
				"username": username,
				"flv_url":  url,
			})
			return
		}
	}
	c.String(http.StatusOK, "room no found!")
}

func myLive(c *gin.Context) {
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	roomPath := utils.CreateRoomPath(username.(string))
	c.HTML(http.StatusOK, "my_live.tmpl", gin.H{
		"username": username,
		"webrtc_url": fmt.Sprintf("webrtc://%s:%s%s",
			setting.SrsSetting.Ip, setting.SrsSetting.HttpApiPort,
			roomPath),
		"rtmp_url": fmt.Sprintf("rtmp://%s:%s%s",
			setting.SrsSetting.Ip, setting.SrsSetting.RtmpPort,
			roomPath),
	})
}
