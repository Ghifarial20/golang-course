package main

import (
	"rest-api-golang/config"
	"rest-api-golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.POST("/person", inDB.CreatePerson)
	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	router.Run(":5000")
}
