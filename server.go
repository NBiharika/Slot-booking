package main

import (
	"Slot_booking/api"
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func main() {
	start_up.Initialize() //
	//manager.Createdatabase()
	server := gin.Default()
	//Health-check
	server.GET("/api/health-check", api.HealthCheck)

	//Slot
	server.GET("/api/slot", api.GetSlot)

	server.POST("/api/add-slot", api.AddSlot)

	//User
	server.GET("/api/user", api.GetUser)

	server.POST("/api/add-user", api.AddUser)

	//Booking
	server.GET("/api/booking", api.GetBooking)

	server.POST("/api/add-booking", api.BookSlot)

	server.PUT("/api/cancel-booking", api.CancelBooking)

	//Get all slots for a user
	server.GET("/api/user-slots", api.UserSlot)

	server.Run(":8080")
}
