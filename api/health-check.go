package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(ctx *gin.Context) {
	//query := ctx.Query("users")
	//path := ctx.Param("users")
	//body := ctx.Request.Body("users")
	err := start_up.BookingController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "there is error"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "okk!!"})
	}
}
