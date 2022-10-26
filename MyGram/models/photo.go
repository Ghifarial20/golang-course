package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Photo struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `json:"title" form:"title" valid:"required~Title is required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoUrl  string    `json:"photo_url" form:"photo_url" valid:"required~photo_url is required"`
	UserID    uint64    `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"User" valid:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p Photo) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(p)

	return result, err
}
