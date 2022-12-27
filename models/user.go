package models

import (
	"fmt"

	"lucy/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                 string `gorm:"name"`
	Salt                 string `gorm:"salt"`
	AuthenticationString string `gorm:"authentication_string"`
}

func (u *User) getAuthenticationString(password string) string {
	return utils.ToMd5(fmt.Sprintf("%s%s", password, u.Salt))
}

func CheckAuth(username string, password string) (bool, error) {
	var user User
	err := db.Where(User{Name: username}).First(&user).Error

	if err != nil {
		return false, err
	}

	if user.AuthenticationString == utils.ToMd5(password+user.Salt) {
		return true, nil
	}

	return false, nil
}

func IsUserExisted(username string) bool {
	var user User
	err := db.Where(User{Name: username}).First(&user).Error

	return err == nil
}

func CreateUser(username string, password string) {
	user := &User{}
	user.Name = username
	user.Salt = utils.RandStr(32)
	user.AuthenticationString = user.getAuthenticationString(password)

	db.Create(user)
}
