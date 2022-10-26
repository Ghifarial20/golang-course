package models

import (
	"errors"
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex" json:"username" valid:"required~Username is required" form:"username"`
	Email     string    `gorm:"uniqueIndex" json:"email" valid:"email~Invalid email format,required~Email is required" form:"email"`
	Password  string    `json:"password"  valid:"required~Username is required,minstringlength(6)~Password at least should be 6" form:"password"`
	Age       uint64    `json:"age" valid:"required~Age is required" form:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user User) Validate() (bool, error) {
	if user.Age < 8 {
		return false, errors.New("Age need to be 8+")
	}

	result, err := govalidator.ValidateStruct(user)
	return result, err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var user User

	errCreate := tx.Model(&User{}).Where("username = ?", u.Username).First(&user).Error
	if errCreate == nil {
		err = errors.New("username already used")
	}

	errCreate = tx.Model(&User{}).Where("email = ?", u.Email).First(&user).Error
	if errCreate == nil {
		err = errors.New("email already used")
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
