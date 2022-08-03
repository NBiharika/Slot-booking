package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func UserSlot(ctx *gin.Context) {
	getUserSlots := GetUserBookedSlots(ctx)

	ctx.HTML(http.StatusOK, "userBookedSlots.html", gin.H{
		"title":     "userSlots",
		"userSlots": getUserSlots,
	})
}

func GetUserBookedSlots(ctx *gin.Context) map[string]map[uint64]interface{} {
	userSlots, _ := start_up.BookingController.GetUserSlot(ctx)
	loc, _ := time.LoadLocation("Asia/Kolkata")

	m := make(map[string]map[uint64]interface{})
	for _, userSlot := range userSlots {
		if m[userSlot.Date] == nil {
			m[userSlot.Date] = make(map[uint64]interface{})
		}
		dateStr := userSlot.Date + " " + userSlot.StartTime
		slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)

		if slotDate.After(time.Now()) {
			m[userSlot.Date][userSlot.ID] = map[string]interface{}{
				"startTime": userSlot.StartTime,
				"status":    "booked",
			}
		}
	}
	return m
}
