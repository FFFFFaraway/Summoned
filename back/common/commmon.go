package common

import (
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
	UserID uint
	Type   string `form:"type" json:"type"`
	Name   string `form:"name" json:"name"`
	Desc   string `form:"desc" json:"desc"`
	People int    `form:"people" json:"people"`
	Ddl    string `form:"ddl" json:"ddl"`
	Img    string
	Status string
}

type Request struct {
	gorm.Model
	SummonedID uint
	UserID     uint
	Desc       string `form:"desc" json:"desc"`
	// Status: ["Not", "Waiting", "Accepted", "Rejected"]
	Status string
}

type Transaction struct {
	gorm.Model
	OwnerID   uint
	TakerID   uint
	OwnerCost int
	TakerCost int
}

var DB *gorm.DB
