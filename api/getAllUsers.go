package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(ctx *gin.Context) {
	users, _ := start_up.UserController.GetAllUsers()

	ctx.HTML(http.StatusOK, "allUsers.html", gin.H{
		"title": "AllUsers",
		"users": users,
	})
}
