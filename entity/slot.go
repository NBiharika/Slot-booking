package entity

import (
	//"fmt"

	"fmt"
	"math"
	"time"
)

type Slot struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Date      string    `json:"date" gorm:"type:varchar(16)"`
	StartTime string    `json:"start_time" gorm:"<-:create"`
	CreatedOn time.Time `json:"created_on" gorm:"autoUpdateTime:milli"`
	UpdatedOn time.Time `json:"updated_on" gorm:"autoUpdateTime:nano"`
}

func (Slot) TableName() string {
	return "slot"
}

func DateForSlot(date time.Time) string {
	DateFormat := "2006-01-02"
	formattedDate := date.Format(DateFormat)
	return formattedDate
}

func PresentTime() string {
	StartTimeFormat := "15:04"
	now := time.Now()
	formattedTime := now.Format(StartTimeFormat)
	return formattedTime
}

func StartTimeOfSlot(j int) string {
	hour := 10.0
	minute := 0.5
	minute = minute * float64(j)
	startTime := hour + minute
	startTimeH := math.Floor(startTime)
	decimalVal := startTime - startTimeH
	var startTimeM float64
	if decimalVal == .5 {
		startTimeM = 3
	} else {
		startTimeM = 0
	}
	return fmt.Sprintf("%v", startTimeH) + ":" + fmt.Sprintf("%v0", startTimeM)
}
