package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
}

const (
	PasswordCost = 18
)

func (user *User) setPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) checkPassword(password string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return error == nil
}
