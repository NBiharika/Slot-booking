package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
)

func GenerateToken(ctx *gin.Context) {
	tokenString, err, statusCode := start_up.TokenController.GenerateToken(ctx)
	if err != nil {
		ctx.HTML(statusCode, "index.html", gin.H{"error": err.Error()})
		//ctx.JSON(statusCode, gin.H{"error": err.Error()})
	} else {
		ctx.HTML(statusCode, "index.html", gin.H{"token": tokenString})
		//ctx.JSON(statusCode, gin.H{"token": tokenString})
	}
}
