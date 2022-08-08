package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeRoleToUser(ctx *gin.Context) {
	err := start_up.UserController.ChangeRoleToUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Role updated to user successfully"})
	}
}
