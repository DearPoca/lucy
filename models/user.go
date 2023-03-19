package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                 string `gorm:"name"`
	Salt                 string `gorm:"salt"`
	AuthenticationString string `gorm:"authentication_string"`
	Email                string `gorm:"email"`
	Telephone            string `gorm:"telephone"`
}
