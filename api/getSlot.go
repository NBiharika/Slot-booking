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
	finalUserSlots, date := FinalUserSlots(ctx)

	ctx.HTML(http.StatusOK, "slot.html", gin.H{
		"title": "slots",
		"slots": finalUserSlots,
		"date":  date,
	})
	//ctx.JSON(http.StatusOK, start_up.SlotController.FindAll())
}

func FinalUserSlots(ctx *gin.Context) (map[uint64]interface{}, string) {
	slots := start_up.SlotController.FindAll()
	userSlots, _ := start_up.BookingController.GetUserSlot(ctx)

	date := slots[0].Date
	m := make(map[uint64]interface{})
	for i := 0; i < len(slots); i++ {
		slotTimeH, _ := strconv.Atoi(slots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(slots[i].StartTime[3:])
		presentTimeH, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[:2])
		presentTimeM, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[3:])
		if slotTimeH > presentTimeH || (slotTimeH == presentTimeH && slotTimeM >= presentTimeM) {
			m[slots[i].ID] = map[string]interface{}{
				"startTime": slots[i].StartTime,
				"status":    "cancelled",
			}
		} else {
			m[slots[i].ID] = map[string]interface{}{
				"startTime": slots[i].StartTime,
				"status":    "expired",
			}
		}
	}
	for i := 0; i < len(userSlots); i++ {
		slotTime, _ := time.Parse("15:04", userSlots[i].StartTime)
		presentTime, _ := time.Parse("15:04", entity.PresentTimePlus30minutes())

		if presentTime.Before(slotTime) {
			m[userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		}
	}
	return m, date
}
