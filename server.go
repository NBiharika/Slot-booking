package main

import (
	"Slot_booking/api"
	"Slot_booking/controller"
	"Slot_booking/middleware"
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func main() {
	start_up.Initialize()

	server := gin.Default()
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.GET("/api/health-check", api.HealthCheck)

	authApis := server.Group("/api/v1/", middleware.Auth())
	{
		authApis.GET("user", api.GetUser)
		authApis.POST("add-booking", api.BookSlot)
		authApis.PUT("cancel-booking", api.CancelBooking)
		authApis.GET("user-slots", api.UserSlot)
		authApis.GET("get-slot", api.GetSlot)
	}

	server.POST("/api/add-slot", api.AddSlot)

	server.POST("/api/add-user", api.AddUser)

	server.GET("/api/booking", api.GetBooking)

	server.POST("api/generate-token", api.GenerateToken)

	secured := server.Group("/secured").Use(middleware.Auth())
	{
		secured.GET("/ping", controller.Ping)
	}
	server.GET("/", controller.LoginAndRegister)

	server.Run(":8080")
}
