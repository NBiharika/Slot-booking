package api

import (
	"Slot_booking/entity"
	"Slot_booking/start_up"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		slotTimeH, _ := strconv.Atoi(slots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(slots[i].StartTime[3:])
		presentTimeH, _ := strconv.Atoi(entity.PresentTime()[:2])
		presentTimeM, _ := strconv.Atoi(entity.PresentTime()[3:])
		todayDateMonth, _ := strconv.Atoi(startDate[5:7])
		dateMonth, _ := strconv.Atoi(slots[i].Date[5:7])
		todayDateDay, _ := strconv.Atoi(startDate[8:])
		dateDay, _ := strconv.Atoi(slots[i].Date[8:])

		if (todayDateMonth == dateMonth && todayDateDay == dateDay) && (slotTimeH < presentTimeH || (slotTimeH == presentTimeH && slotTimeM <= presentTimeM)) {
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
		slotTimeH, _ := strconv.Atoi(userSlots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(userSlots[i].StartTime[3:])
		presentTimeH, _ := strconv.Atoi(entity.PresentTime()[:2])
		presentTimeM, _ := strconv.Atoi(entity.PresentTime()[3:])
		todayDateMonth, _ := strconv.Atoi(startDate[5:7])
		dateMonth, _ := strconv.Atoi(userSlots[i].Date[5:7])
		todayDateDay, _ := strconv.Atoi(startDate[8:])
		dateDay, _ := strconv.Atoi(userSlots[i].Date[8:])

		if (todayDateMonth == dateMonth && todayDateDay == dateDay) && (presentTimeH < slotTimeH || (presentTimeH == slotTimeH && presentTimeM < slotTimeM)) {
			m[userSlots[i].Date][userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		} else if (todayDateMonth < dateMonth) || (todayDateMonth == dateMonth && todayDateDay < dateDay) {
			m[userSlots[i].Date][userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		}
	}
	return m
}
