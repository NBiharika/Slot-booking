package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HealthCheck(ctx *gin.Context) {

	slotTime, _ := time.Parse("15:04", "20:30")
	fmt.Println(slotTime)
	presentTime := time.Now()
	ist, _ := time.LoadLocation("Asia/Kolkata")
	slotTime = time.Date(presentTime.Year(), presentTime.Month(), presentTime.Day(), 20, 30, 0, 0, ist)
	fmt.Println(slotTime)
	ctx.JSON(http.StatusOK, gin.H{"message": "everything ok"})
}
