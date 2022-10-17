package main

import (
	"ass-03/controllers"
	"ass-03/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.StartDB()
	inDB := &controllers.InDB{DB: db}

	r := gin.Default()

	r.GET("/weather", inDB.GetDataWheather)
	r.PUT("/weather", inDB.UpdateWeather)
	r.Run(":5050")
}
