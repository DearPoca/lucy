package media_service

import (
	"fmt"

	"lucy/utils"
)

const rndTokenLength = 16
const prefix = "/lucy"

type Room struct {
	Id         string `json:"id"`
	Owner      string `json:"owner"`
	WebrtcLink string `json:"webrtc link"`
	FlvLink    string `json:"flv link"`
}

func GetRooms() []Room {
	streams := GetStreams()
	rooms := make([]Room, 0)
	for i, _ := range streams {
		if !streams[i].Publish.Active || !VerifyPath(streams[i].Url) {
			continue
		}
		r := Room{
			Id:         streams[i].Id,
			Owner:      ParseUserFromRoomPath(streams[i].Url),
			WebrtcLink: fmt.Sprintf("/play/webrtc?room_id=%s", streams[i].Id),
			FlvLink:    fmt.Sprintf("/play/flv?room_id=%s", streams[i].Id),
		}
		rooms = append(rooms, r)
	}
	return rooms
}

func GenerateRoomPath(username string) string {
	return fmt.Sprintf("%s/%s/%s", prefix, username, utils.RandStr(rndTokenLength))
}

func ParseUserFromRoomPath(path string) string {
	return path[len(prefix)+1 : len(path)-rndTokenLength-1]
}

func VerifyPath(path string) bool {
	if len(path) < rndTokenLength+len(prefix)+3 {
		return false
	}
	if path[:len(prefix)] != prefix {
		return false
	}
	if path[len(path)-rndTokenLength-1] != '/' {
		return false
	}
	return true
}
