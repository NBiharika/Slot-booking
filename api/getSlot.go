package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSlot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, start_up.SlotController.FindAll())
}
