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

type SocialMediaGet struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name" form:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UserSocialMedia
}

type UserSocialMedia struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (idb *InDB) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = uint64(UserID)
	err := idb.DB.Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func (idb *InDB) GetSocialMedias(c *gin.Context) {
	SocialMedia := []models.SocialMedia{}
	response := []SocialMediaGet{}

	err := idb.DB.Preload("User").Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	if len(SocialMedia) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "Failed",
			"Message": "Data not found",
		})
		return
	}

	for i := 0; i < len(SocialMedia); i++ {
		temp := SocialMediaGet{
			ID:             uint(SocialMedia[i].ID),
			Name:           SocialMedia[i].Name,
			SocialMediaURL: SocialMedia[i].SocialMediaUrl,
			UserID:         uint(SocialMedia[i].UserID),
			CreatedAt:      SocialMedia[i].CreatedAt,
			UpdatedAt:      SocialMedia[i].UpdatedAt,
			User: UserSocialMedia{
				ID:       uint(SocialMedia[i].User.ID),
				Username: SocialMedia[i].User.Username,
			},
		}
		response = append(response, temp)
	}
	c.JSON(http.StatusOK, response)
}

func (idb *InDB) UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}
	newSocialMedia := models.SocialMedia{}
	sosMedId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = uint64(userID)
	SocialMedia.ID = uint64(sosMedId)

	err := idb.DB.Model(&SocialMedia).Where("id = ?", sosMedId).Updates(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	idb.DB.Where("id = ?", sosMedId).First(&newSocialMedia)

	c.JSON(http.StatusOK, gin.H{
		"id":               newSocialMedia.ID,
		"name":             newSocialMedia.Name,
		"social_media_url": newSocialMedia.SocialMediaUrl,
		"user_id":          newSocialMedia.UserID,
		"update_at":        newSocialMedia.UpdatedAt,
	})
}

func (idb *InDB) DeleteSocialMedia(c *gin.Context) {
	Sosmed := models.SocialMedia{}
	SosmedID := c.Param("socialMediaId")

	err := idb.DB.Where("id = ?", SosmedID).Delete(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been succesfully deleted",
	})
}
