package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginAndRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "LoginAndRegister"})
}
