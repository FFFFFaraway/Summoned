package service

import (
	"back/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func Signup(context *gin.Context) (error, gin.H){
	var user common.User
	err = common.DB.Where("username = ?", context.PostForm("username")).First(&user).Error
	// no duplicate users, sign up
	if err == gorm.ErrRecordNotFound {
		user := common.User{
			Username: context.PostForm("username"),
			Password: context.PostForm("password"),
		}
		common.DB.Create(&user)
		session := sessions.Default(context)
		session.Clear()
		session.Set("user_id", user.ID)
		if err = session.Save(); err != nil {
			return err, nil
		}
		return nil, gin.H{"message": "successfully sign up"}
	}
	return nil, gin.H{"message": "duplicate user"}
}

func Login(context *gin.Context) (error, gin.H){
	var user common.User
	err = common.DB.Where("username = ?", context.PostForm("username")).First(&user).Error
	if err == gorm.ErrRecordNotFound { // user doesn't exist, abort
		return nil, gin.H{"message": "user doesn't exist"}
	}
	if user.Password == context.PostForm("password") {
		session := sessions.Default(context)
		session.Clear()
		session.Set("user_id", user.ID)
		if err = session.Save(); err != nil {
			return err, nil
		}
		return nil, gin.H{
			"message": "successfully log in",
			"user_id": user.ID,
		}
	}
	return nil, gin.H{"message": "wrong password"}
}

func GetLoginUserId(context *gin.Context) int{
	userId := DefaultUserId(context)
	if userId == nil {
		return -1
	}
	val := int(userId.(uint))
	return val
}

func IsLogin(context *gin.Context) bool {
	return GetLoginUserId(context) >= 0
}

func DefaultUserId(context *gin.Context) interface{}{
	session := sessions.Default(context)
	return session.Get("user_id")
}