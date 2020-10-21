package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
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

func signup(context *gin.Context) {
	var user User
	err = db.Where("username = ?", context.PostForm("username")).First(&user).Error
	// no duplicate users, sign up
	if err == gorm.ErrRecordNotFound {
		user := User{
			Username: context.PostForm("username"),
			Password: context.PostForm("password"),
		}
		db.Create(&user)
		session := sessions.Default(context)
		session.Clear()
		session.Set("user_id", user.ID)
		if err = session.Save(); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "successfully sign up",
		})
	} else { // duplicate users, abort
		context.JSON(http.StatusOK, gin.H{
			"message": "duplicate user",
		})
	}
}

func login(context *gin.Context) {
	var user User
	err = db.Where("username = ?", context.PostForm("username")).First(&user).Error
	if err == gorm.ErrRecordNotFound { // user doesn't exist, abort
		context.JSON(http.StatusOK, gin.H{
			"message": "user doesn't exist",
		})
	} else {
		if user.Password == context.PostForm("password") {
			session := sessions.Default(context)
			session.Clear()
			session.Set("user_id", user.ID)
			if err = session.Save(); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"message": "err",
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"message": "successfully log in",
				"user_id": user.ID,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"message": "wrong password",
			})
		}
	}
}

func getProfile(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get("user_id")
	var user User
	db.First(&user, userId)
	context.JSON(http.StatusOK, user)
}

func updateProfile(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get("user_id")
	var user User
	db.First(&user, userId)
	if context.PostForm("password") != "" {
		user.Password = context.PostForm("password")
	}
	if context.PostForm("phone") != "" {
		user.Phone = context.PostForm("phone")
	}
	db.Save(&user)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully updated",
	})
}

func getAllSummoned(context *gin.Context) {
	var summoneds []Summoned
	db.Find(&summoneds)
	context.JSON(http.StatusOK, summoneds)
}

func getSummoned(context *gin.Context) {
	id := context.Param("id")
	var summoned Summoned
	err = db.Where("id = ?", id).First(&summoned).Error
	if err == gorm.ErrRecordNotFound {
		context.JSON(http.StatusOK, gin.H{
			"message": "no summoned whose id == " + id,
		})
		return
	}
	context.JSON(http.StatusOK, summoned)
}

func newSummoned(context *gin.Context) {
	var summoned Summoned
	if err = context.ShouldBind(&summoned); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bind failed"})
		return
	}
	session := sessions.Default(context)
	summoned.UserID = int(session.Get("user_id").(uint))
	summoned.Status = "Waiting"
	file, header, err := context.Request.FormFile("img")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err = ioutil.WriteFile("./img/"+header.Filename, content, 0644); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	summoned.Img = header.Filename
	db.Create(&summoned)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully add summoned",
	})
}

func getSummonedByOP(context *gin.Context, op string) {
	session := sessions.Default(context)
	userId := session.Get("user_id")
	var summoneds []Summoned
	db.Where("user_id " + op + " ?", userId).Find(&summoneds)
	context.JSON(http.StatusOK, summoneds)
}

func getSummonedByDefault(context *gin.Context) {
	getSummonedByOP(context, "=")
}

func updateSummonedByDefault(context *gin.Context) {
	var summoned Summoned
	if err = context.ShouldBind(&summoned); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bind failed"})
		return
	}
	file, header, err := context.Request.FormFile("img")
	var summonedInMysql Summoned
	db.First(&summonedInMysql, summoned.ID)
	summonedInMysql.Type = summoned.Type
	summonedInMysql.Name = summoned.Name
	summonedInMysql.Desc = summoned.Desc
	summonedInMysql.People = summoned.People
	summonedInMysql.Ddl = summoned.Ddl
	if err != nil {
		fmt.Printf("%v\n", err)
		print("image didn't upload, we keep it\n")
		db.Save(&summonedInMysql)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err = ioutil.WriteFile("./img/"+header.Filename, content, 0644); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	summonedInMysql.Img = header.Filename
	db.Save(&summonedInMysql)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully update summoned",
	})
}

func deleteSummonedByDefault(context *gin.Context) {
	var summoned Summoned
	if err = context.ShouldBind(&summoned); err != nil {
		fmt.Printf("%v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bind failed"})
		return
	}
	var summonedInMysql Summoned
	db.First(&summonedInMysql, summoned.ID)
	db.Delete(&summonedInMysql)
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully delete summoned",
	})
}

func getSummonedExceptDefault(context *gin.Context) {
	getSummonedByOP(context, "<>")
}

func getLoginUserId(context *gin.Context) int{
	session := sessions.Default(context)
	userId := session.Get("user_id")
	if userId == nil {
		return -1
	}
	val := int(userId.(uint))
	return val
}

func getIsLogin(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"signed": getLoginUserId(context),
	})
}

func logout(context *gin.Context){
	session := sessions.Default(context)
	session.Set("dummy", "content") // this will mark the session as "written"
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
}

func loginMid(context *gin.Context) {
	if getLoginUserId(context) >= 0 {
		context.Next()
	}else{
		context.JSON(http.StatusUnauthorized, nil)
		context.Abort()
	}
}
