package main

import (
	"Slot_booking/api"
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func main() {
	start_up.Initialize()

	server := gin.Default()

	server.GET("/api/health-check", api.HealthCheck)

	server.GET("/api/slot", api.GetSlot)

	server.POST("/api/add-slot", api.AddSlot)

	server.GET("/api/user", api.GetUser)

	server.POST("/api/add-user", api.AddUser)

	server.GET("/api/booking", api.GetBooking)

	server.POST("/api/add-booking", api.BookSlot)

	server.PUT("/api/cancel-booking", api.CancelBooking)

	server.GET("/api/user-slots", api.UserSlot)

	server.POST("api/generate-token", api.GenerateToken)

	server.Run(":8080")
}
