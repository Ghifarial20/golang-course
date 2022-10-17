package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          uint   `json:"LineItemId" gorm:"primary_key;autoIncrement"`
	ItemCode    string `json:"itemCode" gorm:"not null;"`
	Description string `json:"description" gorm:"not null;"`
	Quantity    int    `json:"quantity" gorm:"not null;"`
	OrderID     uint   `json:"-" gorm:"not null;"`
}
