package service

import (
	"back/common"
)

func GetUserById(userId interface{}) *common.User {
	var user common.User
	common.DB.First(&user, userId)
	return &user
}