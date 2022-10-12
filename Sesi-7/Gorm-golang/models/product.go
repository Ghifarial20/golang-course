package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product BeforeCreate()")

	if len(p.Name) < 4 {
		err = errors.New("Product name to short")
	}
	return

}
