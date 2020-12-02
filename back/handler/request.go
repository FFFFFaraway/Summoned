package handler

import (
	"back/common"
	"back/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetRequestStatus(context *gin.Context) {
	summonedId := context.Param("id")
	userId := service.DefaultUserId(context)
	var req common.Request
	err = common.DB.Where("summoned_id = ? and user_id = ? and status <> 'Cancelled'", summonedId, userId).First(&req).Error
	if err == gorm.ErrRecordNotFound { // haven't sent request yet
		context.JSON(http.StatusOK, gin.H{
			"requestStatus": "Not",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"requestStatus": req.Status,
	})
}

func UpdateRequestStatus(context *gin.Context) {
	summonedId := context.Param("id")
	userId := context.PostForm("user_id")
	var req common.Request
	err = common.DB.Where("summoned_id = ? and user_id = ?", summonedId, userId).First(&req).Error
	if err == gorm.ErrRecordNotFound { // haven't sent request yet
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "This request doesn't exist",
		})
		return
	}
	req.Status = context.PostForm("status")
	common.DB.Save(&req)
	if req.Status == "Accepted" {
		var trans common.Transaction
		trans.OwnerID = uint(service.GetLoginUserId(context))
		var takerID int
		if takerID, err = strconv.Atoi(userId); err != nil {
			fmt.Printf("%v\n", err)
			context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
			return
		}
		trans.TakerID = uint(takerID)
		var people int
		if people, err = strconv.Atoi(context.PostForm("people")); err != nil {
			fmt.Printf("%v\n", err)
			context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
			return
		}
		trans.OwnerCost = 3 * people
		trans.TakerCost = 1
		common.DB.Create(&trans)

		var summoned common.Summoned
		common.DB.First(&summoned, req.SummonedID)
		summoned.Status = "Complete"
		common.DB.Save(&summoned)

		context.JSON(http.StatusOK, nil)
	}
}

func GetRequest(context *gin.Context) {
	summonedId := context.Param("id")
	var reqs []common.Request
	err = common.DB.Where("summoned_id = ?", summonedId).Find(&reqs).Error
	context.JSON(http.StatusOK, gin.H{
		"requests": reqs,
	})
}

func GetRequestByUser(context *gin.Context) {
	userId := service.DefaultUserId(context)
	var reqs []common.Request
	err = common.DB.Where("user_id = ?", userId).Find(&reqs).Error
	context.JSON(http.StatusOK, gin.H{
		"requests": reqs,
	})
}

func NewRequest(context *gin.Context) {
	var req common.Request
	var summonedID int
	if summonedID, err = strconv.Atoi(context.PostForm("id")); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
		return
	}
	req.SummonedID = uint(summonedID)
	req.UserID = service.DefaultUserId(context).(uint)
	req.Desc = context.PostForm("desc")
	req.Status = "Waiting"
	common.DB.Create(&req)
	context.JSON(http.StatusOK, nil)
}

func UpdateRequest(context *gin.Context) {
	var reqId int
	if reqId, err = strconv.Atoi(context.PostForm("id")); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
		return
	}
	var reqInMysql common.Request
	common.DB.First(&reqInMysql, reqId)
	reqInMysql.Desc = context.PostForm("desc")
	common.DB.Save(&reqInMysql)
	context.JSON(http.StatusOK, nil)
}

func DeleteRequest(context *gin.Context) {
	var reqId int
	if reqId, err = strconv.Atoi(context.Param("id")); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
		return
	}
	var reqInMysql common.Request
	common.DB.First(&reqInMysql, reqId)
	reqInMysql.Status = "Cancelled"
	common.DB.Save(&reqInMysql)
	context.JSON(http.StatusOK, nil)
}