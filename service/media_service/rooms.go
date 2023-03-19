package media_service

import (
	"fmt"

	"lucy/utils"
)

const rndTokenLength = 16
const prefix = "/lucy"

func GetRooms() {

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
