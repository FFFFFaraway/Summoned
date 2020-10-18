package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func myRouter() *gin.Engine {
	r = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store), corsConfig())

	r.POST("/signup", signup)
	r.POST("/login", login)
	r.GET("/signed", getIsLogin)
	r.GET("/summoned", getAllSummoned)
	r.GET("/summoned/:id", getSummoned)

	r.Use(loginMid)

	r.GET("/profile", getProfile)
	r.POST("/profile", updateProfile)
	r.POST("/mysummoned", newSummoned)
	r.GET("/mysummoned", getSummonedByDefault)
	r.GET("/othersummoned", getSummonedExceptDefault)
	r.POST("/signed", logout)
	return r
}
