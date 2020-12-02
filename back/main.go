package main

import (
	"back/common"
	"back/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var r *gin.Engine
var err error

func main() {
	if err = service.DBConfig(); err != nil {
		return
	}
	defer common.DB.Close()
	r := myRouter()
	if err = r.Run("127.0.0.1:9999"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
