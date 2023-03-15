package routers

import (
	"fmt"
	"net/http"

	"lucy/srs"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func play(c *gin.Context) {
	streams := srs.GetStreams()
	roomId := c.Query("room_id")
	for i, _ := range streams {
		if streams[i].Id == roomId {
			url := fmt.Sprintf("webrtc://%s%s", "localhost", streams[i].Url)
			c.HTML(http.StatusOK, "play.html", gin.H{
				"webrtc_url": url,
			})
			return
		}
	}
	c.String(http.StatusOK, "room no found!")
}
