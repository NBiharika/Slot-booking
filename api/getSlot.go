package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetSlot(ctx *gin.Context) {
	finalUserSlots, isAdmin := FinalUserSlots(ctx)

	ctx.HTML(http.StatusOK, "slot.html", gin.H{
		"title":   "slots",
		"slots":   finalUserSlots,
		"isAdmin": isAdmin,
	})
}

func FinalUserSlots(ctx *gin.Context) (map[string]map[uint64]interface{}, bool) {
	todayTime := time.Now()
	endTime := todayTime.Add(6 * 24 * time.Hour)
	loc, _ := time.LoadLocation("Asia/Kolkata")

	slots := start_up.SlotController.FindAll(ctx, todayTime, endTime)
	isAdmin, userSlots, _ := start_up.BookingController.GetUserSlot(ctx)

	m := make(map[string]map[uint64]interface{})

	for _, slot := range slots {
		if m[slot.Date] == nil {
			m[slot.Date] = make(map[uint64]interface{})
		}

		dateStr := slot.Date + " " + slot.StartTime
		slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)

		if slotDate.Before(todayTime) {
			m[slot.Date][slot.ID] = map[string]interface{}{
				"startTime": slot.StartTime,
				"status":    "expired",
			}
		} else {
			m[slot.Date][slot.ID] = map[string]interface{}{
				"startTime": slot.StartTime,
				"status":    "cancelled",
			}
		}
	}
	for _, userSlot := range userSlots {
		dateStr := userSlot.Date + " " + userSlot.StartTime
		slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)
		if slotDate.After(todayTime) {
			m[userSlot.Date][userSlot.ID] = map[string]interface{}{
				"startTime": userSlot.StartTime,
				"status":    "booked",
			}
		}
	}
	return m, isAdmin
}
