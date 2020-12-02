package handler

import (
	"back/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error

func Signup(context *gin.Context) {
	err, res := service.Signup(context)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, res)
}

func Login(context *gin.Context) {
	err, res := service.Login(context)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, res)
}

func GetIsLogin(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"signed": service.GetLoginUserId(context),
	})
}

func LoginMid(context *gin.Context) {
	if service.IsLogin(context) {
		context.Next()
	}else{
		context.JSON(http.StatusUnauthorized, nil)
		context.Abort()
	}
}

func Logout(context *gin.Context){
	session := sessions.Default(context)
	session.Set("dummy", "content") // this will mark the session as "written"
	session.Options(sessions.Options{MaxAge: -1})
	if err = session.Save(); err != nil {
		fmt.Printf("%v\n", err)
	}
}
