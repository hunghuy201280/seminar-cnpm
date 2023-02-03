package main

import (
	"auth-service/component"
	"auth-service/component/tokenprovider"
	"auth-service/middleware"
	"auth-service/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RunService(db *gorm.DB, port string, provider tokenprovider.Provider, refreshProvider tokenprovider.Provider) error {

	appCtx := component.NewAppContext(db, provider, refreshProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	//r.Use(middleware.Authenticate(appCtx))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	},
	)

	//temp := r.Group("/test")
	//temp.Use(middleware.Authenticate(appCtx))
	//temp.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, common.JS{
	//		"alo": 123,
	//	})
	//})
	ginuser.RegisterHandler(r, appCtx)

	return r.Run(":" + port)
}
