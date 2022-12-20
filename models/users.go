package models

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"max=50"`
	Password string `json:"password" validate:"min=6"`
	Email    string `json:"email" gorm:"unique"`
}

func ValidateFields(user *User) error {
	err := validator.Validate(user)
	if err != nil {
		return err
	}
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) ComparePassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))

	if err != nil {
		return err
	}

	return nil
}
