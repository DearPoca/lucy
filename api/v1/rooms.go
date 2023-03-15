package v1

import (
	"log"
	"net/http"

	"lucy/srs"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	streams := srs.GetStreams()
	log.Printf("Streams: %v", streams)
	c.JSON(http.StatusOK, streams)
}
