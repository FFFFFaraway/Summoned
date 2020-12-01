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
	r.Static("img", "img")

	r.Use(loginMid)

	{
		r.GET("/profile/:id", getProfile)
		r.POST("/profile", updateProfile)
	}
	{
		r.GET("/mysummoned", getSummonedByDefault)
		r.POST("/mysummoned", newSummoned)
		r.PUT("/mysummoned", updateSummonedByDefault)
		r.DELETE("/mysummoned/:id", deleteSummoned)
	}
	{
		r.GET("/requestStatus/:id", getRequestStatus)
		r.PUT("/requestStatus/:id", updateRequestStatus)

		r.GET("/request/:id", getRequest)
		r.GET("/requestByUser", getRequestByUser)
		r.POST("/request", newRequest)
		r.PUT("/request", updateRequest)
		r.DELETE("/request/:id", deleteRequest)
	}
	r.GET("/othersummoned", getSummonedExceptDefault)
	r.POST("/signed", logout)
	r.Use(adminMid)
	r.GET("/users", getAllUsers)
	r.GET("/requestsAll", getAllRequests)
	r.GET("/transaction", getTransactions)
	return r
}
