package main

import (
	"log"
	"os"

	"lucy/pkg/setting"
	"lucy/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	gin.SetMode(gin.DebugMode)

	err := os.MkdirAll(setting.AppSetting.AppRoot, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	routers.Run()
}
