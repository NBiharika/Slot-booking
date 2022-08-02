package api

import (
	"Slot_booking/start_up"
	"Slot_booking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSlot(ctx *gin.Context) {
	m, err := utils.ReadRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err = start_up.SlotController.AddSlot(ctx, m)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "slot added successfully"})
	}
}
