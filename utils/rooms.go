package utils

import "fmt"

const rndTokenLength = 16
const prefix = "/live"

func CreateRoomPath(username string) string {
	return fmt.Sprintf("%s/%s/%s", prefix, username, RandStr(rndTokenLength))
}

func ParseUserFromRoomPath(path string) string {
	return path[len(prefix)+1 : len(path)-rndTokenLength-1]
}
