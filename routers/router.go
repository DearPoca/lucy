package routers

import (
	"fmt"

	"lucy/middleware/jwt"
	"lucy/pkg/setting"
	"lucy/routers/api"
	v1 "lucy/routers/api/v1"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	r.GET("/auth", api.GetAuth)
	r.POST("/register", api.Register)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT)
	{
		apiV1.GET("/userinfo", v1.GetUserInfo)
	}
}

func Run() {
	r.Run(fmt.Sprintf(":%d", setting.AppSetting.Port))
}
