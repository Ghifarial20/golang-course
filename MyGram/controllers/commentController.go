package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CommentGet struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserComment
	Photo     PhotoComment
}
type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id`
}

func (idb *InDB) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = uint64(userID)
	err := idb.DB.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}
func (idb *InDB) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	CommentUpdate := models.Comment{}
	CommentID := c.Param("commentId")

	if contentType == appJSON {
		c.ShouldBindJSON(&CommentUpdate)
	} else {
		c.ShouldBind(&CommentUpdate)
	}

	err := idb.DB.Model(&CommentUpdate).Where("id = ?", CommentID).Updates(&CommentUpdate).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	idb.DB.Where("id = ?", CommentID).First(&Comment)

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func (idb *InDB) GetComment(c *gin.Context) {
	Comments := []models.Comment{}

	err := idb.DB.Preload("User").Preload("Photo").Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}
	if len(Comments) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "Failed",
			"Message": "Data not found",
		})
		return
	}
	responses := []CommentGet{}
	for i := 0; i < len(Comments); i++ {
		temp := CommentGet{
			ID:        uint(Comments[i].ID),
			Message:   Comments[i].Message,
			PhotoID:   uint(Comments[i].PhotoId),
			UserID:    uint(Comments[i].UserId),
			CreatedAt: Comments[i].CreatedAt,
			UpdatedAt: Comments[i].UpdatedAt,
			User: UserComment{
				ID:       uint(Comments[i].User.ID),
				Email:    Comments[i].User.Email,
				Username: Comments[i].User.Username,
			},
			Photo: PhotoComment{
				ID:       uint(Comments[i].Photo.ID),
				Title:    Comments[i].Photo.Title,
				Caption:  Comments[i].Photo.Caption,
				PhotoURL: Comments[i].Photo.PhotoUrl,
				UserID:   uint(Comments[i].Photo.UserID),
			},
		}
		responses = append(responses, temp)
	}
	c.JSON(http.StatusOK, responses)
}
func (idb *InDB) DeleteComment(c *gin.Context) {
	Comment := models.Comment{}
	CommentID := c.Param("commentId")

	err := idb.DB.Where("id = ?", CommentID).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Comment has been succesfully deleted",
	})
}
