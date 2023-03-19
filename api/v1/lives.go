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

func GetLives(c *gin.Context) {
	lives := media_service.GetLives()
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, lives))
}

func Record(c *gin.Context) {
	buf := make([]byte, 1024)
	n, err := c.Request.Body.Read(buf)
	request := struct {
		StreamUrl string `json:"stream_url"`
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
	if err = media_service.Record(request.StreamUrl, username.(string)); err != nil {
		log.Warn("Record failed", "body", string(buf), "username", username, "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
	} else {
		log.Debug("record success", "request", request)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
	}
}
