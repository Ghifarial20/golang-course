package main

import (
	"ass-2/database"
	"ass-2/routers"

	_ "ass-2/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	db := database.DBInit()

	r := routers.SetupRoutes(db)
	//call to browser http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
