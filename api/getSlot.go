package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSlot(ctx *gin.Context) {
	//ctx.Redirect(http.StatusMovedPermanently, "/slot")
	slots := start_up.SlotController.FindAll()
	ctx.HTML(http.StatusOK, "slot.html", gin.H{
		"title": "slots",
		"slots": slots,
	})
	//ctx.HTML(http.StatusOK, "slot.html", gin.H{"slots": start_up.SlotController.FindAll()})
	//ctx.JSON(http.StatusOK, start_up.SlotController.FindAll())
}
