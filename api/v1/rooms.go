package v1

import (
	"fmt"
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
	streams := media_service.GetStreams()
	rooms := make([]room, 0)
	for i, _ := range streams {
		if !media_service.VerifyPath(streams[i].Url) {
			continue
		}
		r := room{
			Id:         streams[i].Id,
			Owner:      media_service.ParseUserFromRoomPath(streams[i].Url),
			WebrtcLink: fmt.Sprintf("/play/webrtc?room_id=%s", streams[i].Id),
			FlvLink:    fmt.Sprintf("/play/flv?room_id=%s", streams[i].Id),
		}
		rooms = append(rooms, r)
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, rooms))
}
