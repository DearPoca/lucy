package v1

import (
	"log"
	"net/http"

	"lucy/pkg/respond"
	"lucy/srs"

	"github.com/gin-gonic/gin"
)

type room struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func GetRooms(c *gin.Context) {
	streams := srs.GetStreams()
	log.Printf("Streams: %v", streams)
	rooms := make([]room, 0)
	for i, _ := range streams {
		r := room{
			Id:   streams[i].Id,
			Name: streams[i].Name,
			Path: streams[i].Url,
		}
		rooms = append(rooms, r)
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess, rooms))
}
