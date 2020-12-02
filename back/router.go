package main

import (
	"back/handler"
	"back/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func myRouter() *gin.Engine {
	r = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store), service.GetCorsConfig())

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)
	r.GET("/signed", handler.GetIsLogin)
	r.GET("/summoned", handler.GetAllSummoned)
	r.GET("/summoned/:id", handler.GetSummoned)
	r.Static("img", "img")

	r.Use(handler.LoginMid)

	{
		r.GET("/profile/:id", handler.GetProfile)
		r.POST("/profile", handler.UpdateProfile)
	}
	{
		r.GET("/mysummoned", handler.GetSummonedByDefaultUser)
		r.POST("/mysummoned", handler.NewSummoned)
		r.PUT("/mysummoned", handler.UpdateSummonedByDefaultUser)
		r.DELETE("/mysummoned/:id", handler.DeleteSummoned)
	}
	{
		r.GET("/requestStatus/:id", handler.GetRequestStatus)
		r.PUT("/requestStatus/:id", handler.UpdateRequestStatus)

		r.GET("/request/:id", handler.GetRequest)
		r.GET("/requestByUser", handler.GetRequestByUser)
		r.POST("/request", handler.NewRequest)
		r.PUT("/request", handler.UpdateRequest)
		r.DELETE("/request/:id", handler.DeleteRequest)
	}
	r.GET("/othersummoned", handler.GetSummonedExceptDefaultUser)
	r.POST("/signed", handler.Logout)
	r.Use(handler.AdminMid)
	r.GET("/users", handler.GetAllUsers)
	r.GET("/requestsAll", handler.GetAllRequests)
	r.GET("/transaction", handler.GetTransactions)
	return r
}
