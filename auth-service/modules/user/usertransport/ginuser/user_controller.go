package ginuser

import (
	"auth-service/component"
	"auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine, ctx component.AppContext) {
	users := engine.Group("/users")

	users.POST("", middleware.Authenticate(ctx), Register(ctx))
	users.GET("/user_info", middleware.Authenticate(ctx), GetUserInfo(ctx))
	users.POST("/login", Login(ctx))
	users.POST("/refresh_token", RefreshToken(ctx))
	users.POST("/update_avatar", middleware.Authenticate(ctx), UpdateAvatar(ctx))
	users.POST("/update_profile", middleware.Authenticate(ctx), UpdateProfile(ctx))
}
