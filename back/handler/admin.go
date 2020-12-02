package handler

import (
	"back/common"
	"back/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMid(context *gin.Context) {
	if service.IsAdmin(context) {
		context.Next()
	}else{
		context.JSON(http.StatusUnauthorized, nil)
		context.Abort()
	}
}

func GetAllUsers(context *gin.Context) {
	var users []common.User
	common.DB.Find(&users)
	context.JSON(http.StatusOK, users)
}

func GetAllRequests(context *gin.Context) {
	var reqs []common.Request
	common.DB.Find(&reqs)
	context.JSON(http.StatusOK, reqs)
}

func GetTransactions(context *gin.Context) {
	var trans []common.Transaction
	common.DB.Find(&trans)
	context.JSON(http.StatusOK, trans)
}

func GetProfits(context *gin.Context) {
	var profits []common.Profit
	common.DB.Select("date, city, summoned_type, sum(count) as count, sum(cost) as cost").
		Group("date, city, summoned_type").Find(&profits)
	context.JSON(http.StatusOK, profits)
}
