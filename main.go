package main

import (
	"log"
	"os"

	_ "lucy/models"
	"lucy/pkg/setting"
	"lucy/routers"
)

func main() {
	err := os.MkdirAll(setting.AppSetting.AppRoot, 0777)
	if err != nil {
		log.Fatal(err)
	}

	routers.Run()
}
