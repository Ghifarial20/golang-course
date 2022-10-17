package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint      `json:"-" gorm:"primary_key;autoIncrement"`
	CustomerName string    `json:"customerName" gorm:"not null;"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"autoCreateTime" `
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:ID"`
}
