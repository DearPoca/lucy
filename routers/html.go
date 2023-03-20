package routers

import (
	"fmt"
	"net/http"

	"lucy/middleware/jwt"
	"lucy/pkg/log"
	"lucy/pkg/respond"
	"lucy/pkg/setting"
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
	streams := media_service.GetStreams()
	liveId := c.Query("live_id")
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	for i, _ := range streams {
		if streams[i].Id == liveId {
			url := fmt.Sprintf("webrtc://%s%s", setting.SrsSetting.Ip, streams[i].Url)
			c.HTML(http.StatusOK, "play_webrtc.tmpl", gin.H{
				"username":   username,
				"webrtc_url": url,
			})
			return
		}
	}
	c.String(http.StatusOK, "live no found!")
}

func playFlv(c *gin.Context) {
	streams := media_service.GetStreams()
	liveId := c.Query("live_id")
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	for i, _ := range streams {
		if streams[i].Id == liveId {
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
	c.String(http.StatusOK, "live no found!")
}

func newLive(c *gin.Context) {
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		username = "user"
	}
	live, err := media_service.GenerateLive(username.(string))
	if err != nil {
		log.Warn("Generate live failed", err, err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		c.Abort()
		return
	}
	c.HTML(http.StatusOK, "new_live.tmpl", gin.H{
		"username":   username,
		"webrtc_url": live.WebrtcUrl,
		"rtmp_url":   live.RtmpUrl,
		"flv_url":    live.FlvUrl,
		"live_name":  live.Name,
	})
}
