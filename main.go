package main

import (
	"os"

	"lucy/pkg/log"

	"lucy/pkg/setting"
	"lucy/routers"
	_ "lucy/service/media_service"
)

func main() {
	err := os.MkdirAll(setting.AppSetting.AppRoot, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	routers.Run()
}
