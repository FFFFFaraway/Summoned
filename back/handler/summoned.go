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

func GetAllSummoned(context *gin.Context) {
	summoneds := service.GetAllSummoned()
	context.JSON(http.StatusOK, summoneds)
}

func GetSummoned(context *gin.Context) {
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

func NewSummoned(context *gin.Context) {
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

func GetSummonedByDefaultUser(context *gin.Context) {
	userId := service.DefaultUserId(context)
	summoneds := service.GetSummonedByUserId(userId)
	context.JSON(http.StatusOK, summoneds)
}

func GetSummonedExceptDefaultUser(context *gin.Context) {
	userId := service.DefaultUserId(context)
	summoneds := service.GetSummonedExceptUserId(userId)
	context.JSON(http.StatusOK, summoneds)
}

func UpdateSummonedByDefaultUser(context *gin.Context) {
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

func DeleteSummoned(context *gin.Context) {
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