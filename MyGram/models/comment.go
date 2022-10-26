package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Comment struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId    uint64    `json:"user_id"`
	PhotoId   uint64    `json:"photo_id" form:"photo_id"`
	Message   string    `json:"message" form:"message" valid:"required~Message is required"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"User" valid:"-"`
	Photo     Photo     `gorm:"foreignKey:PhotoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Photo" valid:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Comment) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(c)

	return result, err
}
