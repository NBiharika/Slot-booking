package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func GetBooking(ctx *gin.Context) {
	ctx.JSON(200, start_up.BookingController.FindAll())
}

//
