package routers

import (
	"fmt"
	"net/http"

	"lucy/api"
	"lucy/api/v1"
	"lucy/middleware/jwt"
	"lucy/pkg/setting"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	r = gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Assets
	r.LoadHTMLGlob("assets/*.tmpl")
	r.Static("/assets/js", "assets/js")
	r.StaticFile("/favicon.ico", "assets/favicon.ico")
	r.StaticFS("/images", http.Dir("assets/images"))

	// Front
	r.GET("/", jwt.JWT, index)
	r.GET("/login", login)
	r.GET("/register", register)
	r.GET("/play", play)

	// Background
	r.POST("/api/register", api.Register)
	r.GET("/api/auth", api.Auth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT)
	{
		apiV1.GET("/get_rooms", v1.GetRooms)
		apiV1.GET("/userinfo", v1.GetUserInfo)
	}
}

func Run() {
	r.Run(fmt.Sprintf(":%d", setting.AppSetting.Port))
}
