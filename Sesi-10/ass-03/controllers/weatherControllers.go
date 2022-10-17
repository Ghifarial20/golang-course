package controllers

import (
	"ass-03/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// update data
func (idb *InDB) UpdateWeather(c *gin.Context) {
	var (
		weather    models.Weather
		newWeather models.Weather
	)

	json.NewDecoder(c.Request.Body).Decode(&newWeather)

	err := idb.DB.Where("id = ?", 1).First(&weather).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  err.Error(),
			"message": "data wheather not found",
		})
		return
	}
	err = idb.DB.Model(&weather).Updates(&newWeather).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":  err.Error(),
			"message": "cannot update weather",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": weather,
	})
}

// get data
func (idb *InDB) GetDataWheather(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	weather := models.Weather{}

	err := idb.DB.Where("id = ?", 1).First(&weather).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  err.Error(),
			"message": "data wheather not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"wind":  weather.Wind,
			"water": weather.Water,
		},
	})

}
