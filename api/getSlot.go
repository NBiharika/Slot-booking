package api

import (
	"Slot_booking/entity"
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetSlot(ctx *gin.Context) {
	finalUserSlots := FinalUserSlots(ctx)

	ctx.HTML(http.StatusOK, "slot.html", gin.H{
		"title": "slots",
		"slots": finalUserSlots,
	})
}

func FinalUserSlots(ctx *gin.Context) map[string]map[uint64]interface{} {
	todayTime := time.Now()
	startDate := entity.DateForSlot(todayTime)
	endTime := todayTime.Add(6 * 24 * time.Hour)
	endDate := entity.DateForSlot(endTime)

	slots := start_up.SlotController.FindAll(startDate, endDate)
	userSlots, _ := start_up.BookingController.GetUserSlot(ctx)

	m := make(map[string]map[uint64]interface{})

	for i := 0; i < len(slots); i++ {
		if m[slots[i].Date] == nil {
			m[slots[i].Date] = make(map[uint64]interface{})
		}
		dateStr := slots[i].Date + " " + slots[i].StartTime
		loc, _ := time.LoadLocation("Asia/Kolkata")
		slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)

		if slotDate.Before(todayTime) {
			m[slots[i].Date][slots[i].ID] = map[string]interface{}{
				"startTime": slots[i].StartTime,
				"status":    "expired",
			}
		} else {
			m[slots[i].Date][slots[i].ID] = map[string]interface{}{
				"startTime": slots[i].StartTime,
				"status":    "cancelled",
			}
		}
	}
	for i := 0; i < len(userSlots); i++ {
		dateStr := userSlots[i].Date + " " + userSlots[i].StartTime
		loc, _ := time.LoadLocation("Asia/Kolkata")
		slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)
		if slotDate.After(todayTime) {
			m[userSlots[i].Date][userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		}
	}
	return m
}
