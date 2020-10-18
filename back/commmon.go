package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Phone    string
}

type Summoned struct {
	gorm.Model
	UserID int
	User   User
	Type   string `form:"type" json:"type"`
	Name   string `form:"name" json:"name"`
	Desc   string `form:"desc" json:"desc"`
	People int    `form:"people" json:"people"`
	Ddl    string `form:"ddl" json:"ddl"`
	Img    string
	Status string
}

var db *gorm.DB
var r *gin.Engine
var err error
