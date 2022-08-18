package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(ctx *gin.Context) {
	AllUsers := AllUsers()

	ctx.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "AllUsers",
		"users": AllUsers,
	})
}

func AllUsers() map[uint64]interface{} {
	users, _ := start_up.UserController.GetAllUsers()

	m := make(map[uint64]interface{})
	for _, user := range users {
		m[user.ID] = map[string]interface{}{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"role":      user.Role,
			"status":    user.Status,
		}
	}
	return m
}
