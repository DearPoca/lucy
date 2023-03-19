package user_service

import (
	"errors"
	"fmt"
	"regexp"

	"lucy/models"
	"lucy/utils"
)

type UserInfo struct {
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Telephone string `json:"telephone,omitempty"`
}

func getAuthenticationString(u *models.User, password string) string {
	return utils.ToMd5(fmt.Sprintf("%s%s", password, u.Salt))
}

func CheckAuth(username string, password string) (bool, error) {
	var u models.User
	err := models.Db().Where(models.User{Name: username}).First(&u).Error

	if err != nil {
		return false, err
	}

	if u.AuthenticationString == utils.ToMd5(password+u.Salt) {
		return true, nil
	}

	return false, nil
}

func IsUserExisted(username string) bool {
	var u models.User
	err := models.Db().Where(models.User{Name: username}).First(&u).Error

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

func GetUserInfo(username string) (*UserInfo, error) {
	var u models.User
	err := models.Db().Where(models.User{Name: username}).First(&u).Error

	if err != nil {
		return nil, err
	}
	ret := &UserInfo{
		Username:  u.Name,
		Email:     u.Email,
		Telephone: u.Telephone,
	}
	return ret, nil
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

	u := &models.User{}
	u.Name = username
	u.Salt = utils.RandStr(32)
	u.AuthenticationString = getAuthenticationString(u, password)
	u.Email = email
	u.Telephone = telephone

	models.Db().Create(u)
	return nil
}
