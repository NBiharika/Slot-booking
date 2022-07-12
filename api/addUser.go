package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(ctx *gin.Context) {
	err, statusCode := start_up.UserController.AddUser(ctx)

	if err != nil {
		ctx.HTML(statusCode, "index.html", gin.H{"error": err.Error()})
		//ctx.JSON(statusCode, gin.H{"error": err.Error()})
	} else {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "User added successfully"})
		//ctx.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
	}
}
