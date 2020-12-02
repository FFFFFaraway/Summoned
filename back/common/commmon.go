package common

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string `form:"username" json:"username"`
	Password   string `form:"password" json:"password"`
	IsAdmin    bool   `form:"is_admin" json:"is_admin"`
	Name       string `form:"name" json:"name"`
	NumberType string `form:"number_type" json:"number_type"`
	Number     string `form:"number" json:"number"`
	Phone      string `form:"phone" json:"phone"`
	// Rank: ["Diamond", "VIP", "Normal"]
	Rank         string `form:"rank" json:"rank"`
	Introduction string `form:"introduction" json:"introduction"`
	City         string `form:"city" json:"city"`
}

type Summoned struct {
	gorm.Model
	UserID uint   `form:"user_id" json:"user_id"`
	Type   string `form:"type" json:"type"`
	Name   string `form:"name" json:"name"`
	Desc   string `form:"desc" json:"desc"`
	People int    `form:"people" json:"people"`
	Ddl    string `form:"ddl" json:"ddl"`
	// img 不通过ShouldBind函数进行绑定
	Img string
	// Status: ["Waiting", "Complete", "Cancelled", "Expired"]
	Status string `form:"status" json:"status"`
}

type Request struct {
	gorm.Model
	SummonedID uint   `form:"summoned_id" json:"summoned_id"`
	UserID     uint   `form:"user_id" json:"user_id"`
	Desc       string `form:"desc" json:"desc"`
	// Status: ["Not", "Cancelled", "Waiting", "Accepted", "Rejected"]
	Status string `form:"status" json:"status"`
}

type Transaction struct {
	gorm.Model
	OwnerID   uint `form:"owner_id" json:"owner_id"`
	TakerID   uint `form:"taker_id" json:"taker_id"`
	OwnerCost int  `form:"owner_cost" json:"owner_cost"`
	TakerCost int  `form:"taker_cost" json:"taker_cost"`
}

type Profit struct {
	Date         string `form:"date" json:"date"`
	City         string `form:"city" json:"city"`
	SummonedType string `form:"summoned_type" json:"summoned_type"`
	Count        uint   `form:"count" json:"count"`
	Cost         int    `form:"cost" json:"cost"`
}

var DB *gorm.DB
