package routers

import (
	"fmt"

	"lucy/controller"
	"lucy/pkg/setting"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.New()
	r.POST("/register", controller.Register)
}

func Run() {
	r.Run(fmt.Sprintf(":%d", setting.AppSetting.Port))
}
