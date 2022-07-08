package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	user, err, statusCode := start_up.UserController.RegisterUser(ctx)

	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(statusCode, gin.H{"user": user})
	}
}
