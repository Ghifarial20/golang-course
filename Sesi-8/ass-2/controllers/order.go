package controllers

import (
	"ass-2/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input payload
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.Orders true "Create order"
// @Success 200 {object} models.Orders
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	// Initiate Variable
	var (
		order  models.Orders
		result gin.H
	)

	// Setup Database
	db := c.MustGet("db").(*gorm.DB)

	// Parsing from request body to order
	json.NewDecoder(c.Request.Body).Decode(&order)

	// Creating to database
	err := db.Model(order).Create(&order).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error creating order data: %s", err.Error()),
		})
		return
	}
	// Response
	result = gin.H{
		"result": order,
	}

	fmt.Println("Create order Success")
	c.JSON(http.StatusCreated, result)
}

// GetOrders godoc
// @Summary Get Details of all orders
// @Description Get Details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Orders
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	var (
		order  []models.Orders
		result gin.H
	)

	db := c.MustGet("db").(*gorm.DB)

	err := db.Preload("Items").Find(&order).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error getting order data: %s", err.Error()),
		})
		return
	}

	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)
}

// GetOrder godoc
// @Summary Get Details order by id
// @Description Get Details order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 200 {object} models.Orders
// @Router /orders/{orderId} [get]
func GetOrder(c *gin.Context) {
	var (
		order  []models.Orders
		result gin.H
	)

	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("orderId")

	err := db.Preload("Items").Where("order_id = ?", id).First(&order).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error getting order data by id: %s", err.Error()),
		})
		return
	}

	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)
}

// UpdateOrder godoc
// @Summary Update data order where orderId
// @Description Update data order where orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 200 {object} models.Orders
// @Router /orders/{orderId} [put]
func UpdateOrder(c *gin.Context) {

	var (
		order    models.Orders
		newOrder models.Orders
		result   gin.H
	)

	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("orderId")

	json.NewDecoder(c.Request.Body).Decode(&order)

	err := db.Preload("Items").Where("order_id = ?", id).First(&order).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Order Data Not Found",
		})
		return
	}

	newOrder.Item = order.Item
	newOrder.CustomerName = order.CustomerName
	newOrder.OrderedAt = order.OrderedAt
	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&newOrder).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error updating order data by id: %s", err.Error()),
		})
		return
	}

	result = gin.H{
		"result": "Update successfully",
	}

	c.JSON(http.StatusOK, result)
}

// DeleteOrder godoc
// @Summary Delete data order where orderId
// @Description Delete data order where orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var (
		order  models.Orders
		result gin.H
	)

	db := c.MustGet("db").(*gorm.DB)

	err := db.First(&order, id).Error

	if err != nil {
		result = gin.H{
			"result": "Data Not Found",
		}
	}

	err = db.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Delete Unsuccessful",
		}
	}
	result = gin.H{
		"result": "Delete successful",
	}

	c.JSON(http.StatusOK, result)
}
