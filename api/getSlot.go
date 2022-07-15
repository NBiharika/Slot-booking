package api

import (
	"Slot_booking/entity"
	"Slot_booking/start_up"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSlot(ctx *gin.Context) {
	finalUserSlots := FinalUserSlots(ctx)

	ctx.HTML(http.StatusOK, "slot.html", gin.H{
		"title": "slots",
		"slots": finalUserSlots,
	})
	//ctx.JSON(http.StatusOK, start_up.SlotController.FindAll())
}

func FinalUserSlots(ctx *gin.Context) map[uint64]interface{} {
	slots := start_up.SlotController.FindAll()
	userSlots, _ := start_up.BookingController.GetUserSlot(ctx)

	m := make(map[uint64]interface{})
	for i := 0; i < len(slots); i++ {
		slotTimeH, _ := strconv.Atoi(slots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(slots[i].StartTime[3:])
		presentTimeH, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[:2])
		presentTimeM, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[3:])
		dateForSlot, _ := strconv.Atoi(slots[i].Date[8:])
		dateToday, _ := strconv.Atoi(entity.DateForSlot()[8:])
		if dateForSlot < dateToday {
			m[slots[i].ID] = map[string]interface{}{
				"date":      slots[i].Date,
				"startTime": slots[i].StartTime,
				"status":    "expired",
			}
		} else if slotTimeH > presentTimeH || (slotTimeH == presentTimeH && slotTimeM >= presentTimeM) || (dateToday >= dateForSlot) {
			m[slots[i].ID] = map[string]interface{}{
				"date":      slots[i].Date,
				"startTime": slots[i].StartTime,
				"status":    "cancelled",
			}
		} else {
			m[slots[i].ID] = map[string]interface{}{
				"date":      slots[i].Date,
				"startTime": slots[i].StartTime,
				"status":    "expired",
			}
		}
	}
	for i := 0; i < len(userSlots); i++ {
		m[userSlots[i].ID] = map[string]interface{}{
			"date":      slots[i].Date,
			"startTime": slots[i].StartTime,
			"status":    "booked",
		}
	}
	fmt.Println("slots", m)
	return m
}
