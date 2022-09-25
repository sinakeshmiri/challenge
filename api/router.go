package api

import "github.com/gin-gonic/gin"

var api apiHandler

func ASCIIArtRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/image", api.RegisterImage)
	router.POST("/image/:sha256/chunks", api.UploadImageChunk)
	router.GET("/image/:sha256", api.DownloadImage)

	return router
}
