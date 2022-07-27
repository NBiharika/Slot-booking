package api

import (
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		slotTimeH, _ := strconv.Atoi(userSlots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(userSlots[i].StartTime[3:])
		dateYear, _ := strconv.Atoi(userSlots[i].Date[0:4])
		dateMonth, _ := strconv.Atoi(userSlots[i].Date[5:7])
		dateDay, _ := strconv.Atoi(userSlots[i].Date[8:])
		slotTime := time.Date(dateYear, time.Month(dateMonth), dateDay, slotTimeH, slotTimeM, 0, 0, time.Local)

		if slotTime.After(time.Now()) {
			m[userSlots[i].Date][userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		}
	}
	return m
}
