package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotosGet struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserPhoto
}

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (idb *InDB) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint64(UserID)
	err := idb.DB.Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})

}

func (idb *InDB) GetPhotos(c *gin.Context) {
	Photos := []models.Photo{}
	response := []PhotosGet{}

	err := idb.DB.Preload("User").Find(&Photos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}
	if len(Photos) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "Failed",
			"Message": "Data not found",
		})
		return
	}

	for i := 0; i < len(Photos); i++ {
		temp := PhotosGet{
			ID:        uint(Photos[i].ID),
			Title:     Photos[i].Title,
			Caption:   Photos[i].Caption,
			PhotoURL:  Photos[i].PhotoUrl,
			UserID:    uint(Photos[i].UserID),
			CreatedAt: Photos[i].CreatedAt,
			UpdatedAt: Photos[i].UpdatedAt,
			User: UserPhoto{
				Email:    Photos[i].User.Email,
				Username: Photos[i].User.Username,
			},
		}
		response = append(response, temp)
	}
	c.JSON(http.StatusOK, response)

}

func (idb *InDB) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	newPhoto := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint64(userID)
	Photo.ID = uint64(photoId)

	err := idb.DB.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	idb.DB.Where("id = ?", photoId).First(&newPhoto)

	c.JSON(http.StatusOK, gin.H{
		"id":         newPhoto.ID,
		"title":      newPhoto.Title,
		"caption":    newPhoto.Caption,
		"photo_url":  newPhoto.PhotoUrl,
		"user_id":    newPhoto.UserID,
		"updated_at": newPhoto.UpdatedAt,
	})
}

func (idb *InDB) DeletePhoto(c *gin.Context) {
	Photo := models.Photo{}
	photoId := c.Param("photoId")

	err := idb.DB.Where("id =?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo has been succesfully deleted",
	})
}
