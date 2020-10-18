package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err = gorm.Open("mysql", "root:123456@/summoned?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Summoned{})
	defer db.Close()
	r := myRouter()
	if err = r.Run("127.0.0.1:9999"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
