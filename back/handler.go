package main

import (
	"back/common"
	"back/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
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

func updateProfile(context *gin.Context) {
	userId := service.DefaultUserId(context)
	user, _ := service.GetUserById(userId)
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
	file, _, err := context.Request.FormFile("img")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	userId := service.DefaultUserId(context)
	if err = service.NewSummoned(summoned, file, userId); err != nil {
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
	file, _, err := context.Request.FormFile("img")
	// if no image uploaded for updating, we keep it constant
	keepImg := err != nil
	if err = service.UpdateSummoned(summoned, file, keepImg); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully update summoned",
	})
}

func deleteSummoned(context *gin.Context) {
	var summonedId int
	if summonedId, err = strconv.Atoi(context.Param("id")); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
		return
	}
	if err = service.DeleteSummoned(summonedId); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	context.JSON(http.StatusOK, nil)
}

func getRequestStatus(context *gin.Context) {
	summonedId := context.Param("id")
	userId := service.DefaultUserId(context)
	var req common.Request
	err = common.DB.Where("summoned_id = ? and user_id = ?", summonedId, userId).First(&req).Error
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

func updateRequestStatus(context *gin.Context) {
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
		context.JSON(http.StatusOK, nil)
	}
}

func getRequest(context *gin.Context) {
	summonedId := context.Param("id")
	var reqs []common.Request
	err = common.DB.Where("summoned_id = ?", summonedId).Find(&reqs).Error
	context.JSON(http.StatusOK, gin.H{
		"requests": reqs,
	})
}

func getRequestByUser(context *gin.Context) {
	userId := service.DefaultUserId(context)
	var reqs []common.Request
	err = common.DB.Where("user_id = ?", userId).Find(&reqs).Error
	context.JSON(http.StatusOK, gin.H{
		"requests": reqs,
	})
}

func newRequest(context *gin.Context) {
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

func updateRequest(context *gin.Context) {
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

func deleteRequest(context *gin.Context) {
	var reqId int
	if reqId, err = strconv.Atoi(context.Param("id")); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID Atoi failed"})
		return
	}
	var reqInMysql common.Request
	common.DB.First(&reqInMysql, reqId)
	common.DB.Delete(&reqInMysql)
	context.JSON(http.StatusOK, nil)
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
	if err = session.Save(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

