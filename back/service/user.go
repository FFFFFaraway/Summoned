package service

import (
	"back/common"
)

func GetUserById(userId interface{}) (*common.User, error) {
	var user common.User
	err = common.DB.First(&user, userId).Error
	return &user, err
}