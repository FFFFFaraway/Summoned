package service

import (
	"back/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func GetCorsConfig() gin.HandlerFunc{
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://10.128.252.137:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})
}

func DBConfig() error {
	common.DB, err = gorm.Open("mysql", "root:123456@/summoned?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	common.DB.AutoMigrate(&common.User{})
	common.DB.AutoMigrate(&common.Summoned{})
	common.DB.AutoMigrate(&common.Request{})
	common.DB.AutoMigrate(&common.Transaction{})
	common.DB.AutoMigrate(&common.Profit{})
	common.DB.Model(&common.Summoned{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	common.DB.Model(&common.Request{}).
		AddForeignKey("summoned_id", "summoneds(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	common.DB.Model(&common.Transaction{}).
		AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("taker_id", "users(id)", "RESTRICT", "RESTRICT")
	return nil
}