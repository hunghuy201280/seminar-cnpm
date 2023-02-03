package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"s3-service/component"
	"s3-service/component/tokenprovider"
	"s3-service/component/uploadprovider"
	"s3-service/middleware"
	"s3-service/modules/upload/uploadtransportation/ginupload"
)

func RunService(
	db *gorm.DB,
	port string,
	provider tokenprovider.Provider,
	refreshProvider tokenprovider.Provider,
	uploadProvider uploadprovider.UploadProvider,
) error {

	appCtx := component.NewAppContext(db, provider, refreshProvider, uploadProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.Authenticate(appCtx))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	},
	)

	r.POST("/upload", ginupload.Upload(appCtx))

	return r.Run(":" + port)
}
