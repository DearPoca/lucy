package routers

import (
	"fmt"

	"lucy/api"
	v1 "lucy/api/v1"
	"lucy/middleware/jwt"
	"lucy/pkg/setting"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)
	r.POST("/register", api.Register)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT)
	{
		apiV1.GET("/userinfo", v1.GetUserInfo)
		apiV1.POST("/create_bucket", v1.CreateBucket)
	}
}

func Run() {
	r.Run(fmt.Sprintf(":%d", setting.AppSetting.Port))
}
