package v1

import (
	"net/http"

	"lucy/pkg/respond"
	"lucy/service/media_service"

	"github.com/gin-gonic/gin"
)

type room struct {
	Id         string `json:"id"`
	Owner      string `json:"owner"`
	WebrtcLink string `json:"webrtc link"`
	FlvLink    string `json:"flv link"`
}

func GetRooms(c *gin.Context) {
	rooms := media_service.GetRooms()
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, rooms))
}
