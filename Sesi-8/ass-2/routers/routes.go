package routers

import (
	"ass-2/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	r.GET("/orders/:orderId", controllers.GetOrder)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders", controllers.DeleteOrder)
	return r
}
