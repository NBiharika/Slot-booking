package api

import (
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
	//ctx.JSON(http.StatusOK, gin.H{"message": "check"})
}

func FinalUserSlots(ctx *gin.Context) (map[uint64]interface{}, string) {
	//startDate - entity.dateforslot(time.Now)
	//endDate - entity.dateforslot()
	//time.Now().Add(7 * 24 * time.Hour)
	slots := start_up.SlotController.FindAll()
	userSlots, _ := start_up.BookingController.GetUserSlot(ctx)

	date := slots[0].Date
	//new map : map[string]map[uint64]interface{}
	//
	//map["2022-07-22"]=m;

	m := make(map[uint64]interface{})
	for i := 0; i < len(slots); i++ {

		//slotTime, _ := time.Parse("15:04", slots[i].StartTime)
		slotTimeH, _ := strconv.Atoi(slots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(slots[i].StartTime[3:])
		//presentTimePlus30minutes, _ := time.Parse("15:04", entity.PresentTimePlus30minutes())
		presentTime := time.Now()
		ist, _ := time.LoadLocation("Asia/Kolkata")
		slotTime := time.Date(presentTime.Year(), presentTime.Month(), presentTime.Day(), slotTimeH, slotTimeM, 0, 0, ist)
		if presentTime.Before(slotTime) {
			//mp[slots[i].date][slots[i].ID]=map[string]interface{}.....
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
		slotTimeH, _ := strconv.Atoi(userSlots[i].StartTime[:2])
		slotTimeM, _ := strconv.Atoi(userSlots[i].StartTime[3:])

		presentTime := time.Now()
		ist, _ := time.LoadLocation("Asia/Kolkata")
		slotTime := time.Date(presentTime.Year(), presentTime.Month(), presentTime.Day(), slotTimeH, slotTimeM, 0, 0, ist)

		if presentTime.Before(slotTime) {
			m[userSlots[i].ID] = map[string]interface{}{
				"startTime": userSlots[i].StartTime,
				"status":    "booked",
			}
		}
	}
	return m, date
}
