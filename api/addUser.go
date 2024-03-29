package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(ctx *gin.Context) {
	err, statusCode := start_up.UserController.AddUser(ctx)

	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
	}
}
