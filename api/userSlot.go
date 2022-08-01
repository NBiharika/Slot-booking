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

	m := make(map[string]map[uint64]interface{})
	for i := 0; i < len(userSlots); i++ {
		if m[userSlots[i].Date] == nil {
			m[userSlots[i].Date] = make(map[uint64]interface{})
		}
		dateStr := userSlots[i].Date + " " + userSlots[i].StartTime
		loc, _ := time.LoadLocation("Asia/Kolkata")
		slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)

		if slotDate.After(time.Now()) {
			m[userSlots[i].Date][userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		}
	}
	return m
}
