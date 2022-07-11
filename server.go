package main

import (
	"Slot_booking/api"
	"Slot_booking/controller"
	"Slot_booking/middlewares"
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func main() {
	start_up.Initialize()

	server := gin.Default()

	server.GET("/api/health-check", api.HealthCheck)

	authApis := server.Group("/api/v1/", middlewares.Auth())
	{
		authApis.GET("user", api.GetUser)
		authApis.POST("add-booking", api.BookSlot)
		authApis.PUT("cancel-booking", api.CancelBooking)
		authApis.GET("user-slots", api.UserSlot)
	}

	server.GET("/api/slot", api.GetSlot)

	server.POST("/api/add-slot", api.AddSlot)

	server.POST("/api/add-user", api.AddUser)

	server.GET("/api/booking", api.GetBooking)

	server.POST("api/generate-token", api.GenerateToken)

	secured := server.Group("/secured").Use(middlewares.Auth())
	{
		secured.GET("/ping", controller.Ping)
	}

	server.Run(":8080")
}
