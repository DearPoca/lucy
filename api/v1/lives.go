package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lucy/middleware/jwt"
	"lucy/pkg/log"
	"lucy/pkg/respond"
	"lucy/service/media_service"

	"github.com/gin-gonic/gin"
)

func GetActiveLives(c *gin.Context) {
	lives := media_service.GetActiveLives()
	type liveInterface struct {
		Id         string `json:"id"`
		Title      string `json:"title"`
		Owner      string `json:"owner"`
		StartTime  string `json:"start_time"`
		WebrtcLink string `json:"webrtc_link"`
		FlvLink    string `json:"flv_link"`
	}
	var ret []liveInterface
	for i, _ := range lives {
		ret = append(ret, liveInterface{
			Id:         lives[i].Id,
			Title:      lives[i].Title,
			Owner:      lives[i].Owner,
			StartTime:  lives[i].StartTime,
			WebrtcLink: fmt.Sprintf("/play/webrtc?live_id=%s", lives[i].Id),
			FlvLink:    fmt.Sprintf("/play/flv?live_id=%s", lives[i].Id),
		})
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, ret))
}

func RecordLive(c *gin.Context) {
	buf := make([]byte, 1024)
	n, err := c.Request.Body.Read(buf)
	request := struct {
		LiveName string `json:"live_name"`
	}{}
	err = json.Unmarshal(buf[:n], &request)
	if err != nil {
		log.Warn("RecordLive failed", "request", string(buf), "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeParamInvalid))
		return
	}

	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	if err = media_service.LiveRecord(request.LiveName, username.(string)); err != nil {
		log.Warn("RecordLive failed", "body", string(buf), "username", username, "err", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
	} else {
		log.Debug("record success", "request", request)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
	}
}

func ListRecord(c *gin.Context) {
	username := c.Query("username")
	lives, err := media_service.GetLivesByUser(username)
	log.Info("ListRecord", "username", username)
	if err != nil {
		log.Warn("List record failed", "err", err.Error(), "username", username)
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, lives))
}

func NewLive(c *gin.Context) {
	username, exists := c.Get(jwt.KeyOfUsername)
	if !exists {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	title := c.Query("liveTitle")
	if len(title) == 0 {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeLiveTitleEmpty))
		return
	}
	live, err := media_service.GenerateLive(username.(string), title)
	if err != nil {
		log.Warn("Generate live failed", err, err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, live))
}
