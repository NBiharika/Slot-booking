package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	user, err := start_up.UserController.Find(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"user": user})
	}
}
