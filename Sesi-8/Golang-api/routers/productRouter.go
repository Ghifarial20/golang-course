package routers

import (
	"golang-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/albums", controllers.SeedAlbum)
	router.GET("/albums", controllers.ShowAll)
	router.GET("/albums/:title", controllers.ShowSingle)
	router.POST("/getFlexAlbums", controllers.ShowAlbums)
	router.PUT("/albums/:title", controllers.UpdateAlbums)
	router.DELETE("/albums/:title", controllers.DeleteAlbums)

	return router
}
