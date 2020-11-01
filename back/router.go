package main

import (
	"back/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func myRouter() *gin.Engine {
	r = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store), service.GetCorsConfig())

	r.POST("/signup", signup)
	r.POST("/login", login)
	r.GET("/signed", getIsLogin)
	r.GET("/summoned", getAllSummoned)
	r.GET("/summoned/:id", getSummoned)

	r.Use(loginMid)

	{
		r.GET("/profile", getProfile)
		r.POST("/profile", updateProfile)
	}
	{
		r.GET("/mysummoned", getSummonedByDefault)
		r.POST("/mysummoned", newSummoned)
		r.PUT("/mysummoned", updateSummonedByDefault)
		r.DELETE("/mysummoned", deleteSummonedByDefault)
	}
	r.GET("/othersummoned", getSummonedExceptDefault)
	r.POST("/signed", logout)
	return r
}
