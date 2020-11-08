package main

import (
	"back/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var r *gin.Engine
var err error

func main() {
	common.DB, err = gorm.Open("mysql", "root:123456@/summoned?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	common.DB.AutoMigrate(&common.User{})
	common.DB.AutoMigrate(&common.Summoned{})
	common.DB.AutoMigrate(&common.Request{})
	common.DB.AutoMigrate(&common.Transaction{})
	common.DB.Model(&common.Summoned{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	common.DB.Model(&common.Request{}).
		AddForeignKey("summoned_id", "summoneds(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	common.DB.Model(&common.Transaction{}).
		AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("taker_id", "users(id)", "RESTRICT", "RESTRICT")
	defer common.DB.Close()
	r := myRouter()
	if err = r.Run("127.0.0.1:9999"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
