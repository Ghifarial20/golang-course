package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func (idb *InDB) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := idb.DB.Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func (idb *InDB) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := idb.DB.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid email",
			"error":   "Unauthorized",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password salah",
			"error":   "Unauthorized",
		})
		return
	}

	token := helpers.GenerateToken(uint(User.ID), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (idb *InDB) UpdateUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}
	UpdateUser := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := idb.DB.Model(&User).Where("id = ?", userID).Updates(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	idb.DB.Where("id = ?", userID).First(&UpdateUser)

	c.JSON(http.StatusOK, gin.H{
		"id":         userID,
		"email":      UpdateUser.Email,
		"username":   UpdateUser.Username,
		"age":        UpdateUser.Age,
		"updated_at": UpdateUser.UpdatedAt,
	})
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	User := models.User{}
	userID := uint(userData["id"].(float64))

	err := idb.DB.Where("id = ?", userID).Delete(&User).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been succesfully deleted",
	})
}
