package routers

import (
	"fmt"
	"net/http"

	"lucy/api"
	"lucy/api/v1"
	"lucy/middleware/jwt"
	"lucy/pkg/setting"
	"lucy/service/media_service"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	r = gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Assets
	r.LoadHTMLGlob("assets/html/*.tmpl")
	r.StaticFS("/assets/js", http.Dir("assets/js"))
	r.StaticFS("/assets/css", http.Dir("assets/css"))
	r.StaticFS(media_service.LiveRecordPath, http.Dir(
		fmt.Sprintf("%s%s", setting.AppSetting.AppRoot, media_service.LiveRecordPath),
	))
	r.StaticFile("/favicon.ico", "assets/favicon.ico")

	// Front
	r.GET("/login", login)
	r.GET("/register", register)
	r.GET("/", jwt.JWT, index)
	r.GET("/play/webrtc", jwt.JWT, playWebrtc)
	r.GET("/play/flv", jwt.JWT, playFlv)
	r.GET("/userinfo", jwt.JWT, userinfo)
	r.GET("/new_live", jwt.JWT, newLive)

	// Background
	r.POST("/api/register", api.Register)
	r.GET("/api/auth", api.Auth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT)
	{
		apiV1.GET("/list_lives", v1.GetActiveLives)
		apiV1.GET("/userinfo", v1.GetUserInfo)
		apiV1.POST("/record_live", v1.RecordLive)
		apiV1.GET("/list_record", v1.ListRecord)
		apiV1.GET("/new_live", v1.NewLive)
	}
}

func Run() {
	r.Run(fmt.Sprintf(":%d", setting.AppSetting.Port))
}
