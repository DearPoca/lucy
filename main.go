package main

import (
	"log"
	"os"

	"lucy/pkg/setting"
	"lucy/routers"
	_ "lucy/srs"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err := os.MkdirAll(setting.AppSetting.AppRoot, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	routers.Run()
}
