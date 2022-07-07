package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSlot(ctx *gin.Context) {
	slots, err := start_up.BookingController.GetUserSlot(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Slots": slots})
	}
}

//
