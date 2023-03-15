package user_service

import (
	"fmt"

	"lucy/models"
	"lucy/utils"
)

func getAuthenticationString(u *models.User, password string) string {
	return utils.ToMd5(fmt.Sprintf("%s%s", password, u.Salt))
}

func CheckAuth(username string, password string) (bool, error) {
	var user models.User
	err := models.Db().Where(models.User{Name: username}).First(&user).Error

	if err != nil {
		return false, err
	}

	if user.AuthenticationString == utils.ToMd5(password+user.Salt) {
		return true, nil
	}

	return false, nil
}

func IsUserExisted(username string) bool {
	var user models.User
	err := models.Db().Where(models.User{Name: username}).First(&user).Error

	return err == nil
}

func CreateUser(username string, password string) {
	user := &models.User{}
	user.Name = username
	user.Salt = utils.RandStr(32)
	user.AuthenticationString = getAuthenticationString(user, password)

	models.Db().Create(user)
}
