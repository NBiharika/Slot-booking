package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func GenerateToken(ctx *gin.Context) {
	tokenString, err, statusCode := start_up.TokenController.GenerateToken(ctx)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(statusCode, gin.H{"token": tokenString})
	}
}
