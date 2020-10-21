package main

import (
	"back/common"
	"back/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func signup(context *gin.Context) {
	err, res := service.Signup(context)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, res)
}

func login(context *gin.Context) {
	err, res := service.Login(context)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, res)
}

func getIsLogin(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"signed": service.GetLoginUserId(context),
	})
}

func getAllSummoned(context *gin.Context) {
	summoneds := service.GetAllSummoned()
	context.JSON(http.StatusOK, summoneds)
}

func getSummoned(context *gin.Context) {
	id := context.Param("id")
	summoned, err := service.GetSummonedById(id)
	if err == gorm.ErrRecordNotFound {
		context.JSON(http.StatusOK, gin.H{
			"message": "no summoned whose id == " + id,
		})
		return
	}
	context.JSON(http.StatusOK, summoned)
}

func loginMid(context *gin.Context) {
	if service.IsLogin(context) {
		context.Next()
	}else{
		context.JSON(http.StatusUnauthorized, nil)
		context.Abort()
	}
}

func getProfile(context *gin.Context) {
	userId := service.DefaultUserId(context)
	user := service.GetUserById(userId)
	context.JSON(http.StatusOK, user)
}

func updateProfile(context *gin.Context) {
	userId := service.DefaultUserId(context)
	user := service.GetUserById(userId)
	if context.PostForm("password") != "" {
		user.Password = context.PostForm("password")
	}
	if context.PostForm("phone") != "" {
		user.Phone = context.PostForm("phone")
	}
	common.DB.Save(&user)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully updated",
	})
}

func newSummoned(context *gin.Context) {
	var summoned common.Summoned
	if err = context.ShouldBind(&summoned); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bind failed"})
		return
	}
	file, header, err := context.Request.FormFile("img")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	userId := service.DefaultUserId(context)
	if err = service.NewSummoned(summoned, file, header, userId); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully add summoned",
	})
}

func getSummonedByDefault(context *gin.Context) {
	userId := service.DefaultUserId(context)
	summoneds := service.GetSummonedByUserId(userId)
	context.JSON(http.StatusOK, summoneds)
}

func updateSummonedByDefault(context *gin.Context) {
	var summoned common.Summoned
	if err = context.ShouldBind(&summoned); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bind failed"})
		return
	}
	file, header, err := context.Request.FormFile("img")
	// if no image uploaded for updating, we keep it constant
	keepImg := err != nil
	if err = service.UpdateSummoned(summoned, file, header, keepImg); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully update summoned",
	})
}

func deleteSummonedByDefault(context *gin.Context) {
	var summoned common.Summoned
	if err = context.ShouldBind(&summoned); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bind failed"})
		return
	}
	service.DeleteSummoned(summoned)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully delete summoned",
	})
}

func getSummonedExceptDefault(context *gin.Context) {
	userId := service.DefaultUserId(context)
	summoneds := service.GetSummonedExceptUserId(userId)
	context.JSON(http.StatusOK, summoneds)
}

func logout(context *gin.Context){
	session := sessions.Default(context)
	session.Set("dummy", "content") // this will mark the session as "written"
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
}

