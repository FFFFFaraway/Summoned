package handler

import (
	"back/common"
	"back/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProfile(context *gin.Context) {
	var userId int
	if userId, err = strconv.Atoi(context.Param("id")); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
		return
	}
	user, err := service.GetUserById(userId)
	if err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusOK, gin.H{"message": "User doesn't exist"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func UpdateProfile(context *gin.Context) {
	userId := service.DefaultUserId(context)
	user, _ := service.GetUserById(userId)
	if context.PostForm("password") != "" {
		user.Password = context.PostForm("password")
	}
	if context.PostForm("phone") != "" {
		user.Phone = context.PostForm("phone")
	}
	if context.PostForm("introduction") != "" {
		user.Introduction = context.PostForm("introduction")
	}
	common.DB.Save(&user)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully updated",
	})
}