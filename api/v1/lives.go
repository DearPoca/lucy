package v1

import (
	"encoding/json"
	"net/http"

	"lucy/middleware/jwt"
	"lucy/pkg/log"
	"lucy/pkg/respond"
	"lucy/service/media_service"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	rooms := media_service.GetRooms()
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, rooms))
}

func Record(c *gin.Context) {
	buf := make([]byte, 1024)
	n, err := c.Request.Body.Read(buf)
	request := struct {
		RtmpUrl string `json:"rtmp_url"`
	}{}
	err = json.Unmarshal(buf[:n], &request)
	if err != nil {
		log.Warn("Record failed", "request", string(buf), "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeParamInvalid))
		return
	}

	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	if err = media_service.Record(request.RtmpUrl, username.(string)); err != nil {
		log.Warn("Record failed", "body", string(buf), "username", username, "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
	} else {
		log.Debug("record success", "request", request)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
	}
}
