package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func GetSlot(ctx *gin.Context) {
	ctx.JSON(200, start_up.SlotController.FindAll())
}

//
