package user_service

import (
	"errors"
	"fmt"
	"regexp"

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

func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyPhoneFormat(phoneNum string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phoneNum)
}

func CreateUser(username string, password string, email string, telephone string) error {
	if !VerifyPhoneFormat(telephone) {
		return errors.New("telephone number not valid")
	}
	if !VerifyEmailFormat(email) {
		return errors.New("email number not valid")
	}
	if len(username) == 0 || len(password) == 0 {
		return errors.New("username or password not valid")
	}

	user := &models.User{}
	user.Name = username
	user.Salt = utils.RandStr(32)
	user.AuthenticationString = getAuthenticationString(user, password)
	user.Email = email
	user.Telephone = telephone

	models.Db().Create(user)
	return nil
}
