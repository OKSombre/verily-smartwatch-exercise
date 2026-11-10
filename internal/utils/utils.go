package utils

import (
	"fmt"
	"time"
)

type DayState int

const (
	DayStateFuture DayState = iota
	DayStateNoActivity
	DayStateWaitingToSync
	DayStateMet
	DayStateUnmet
)

const MinWearHours = 16
const DateLayout = "2006-01"

type DayStateParam struct {
	Date      string 	// YYYY-MM format
	AsOf      string 	// YYYY-MM format
	WearHours float32 	// Total hours the device was worn
	SyncCount uint 		// Number of times the device was synced
}

func Hello() {
    fmt.Println("Hello from API package!")
}

func GetDayState(param DayStateParam) DayState {
	fmt.Printf("Date: %s, As Of: %s, Wear Hours: %.2f, Sync Count: %d\n", param.Date, param.AsOf, param.WearHours, param.SyncCount)
	
	if param.Date > param.AsOf {
		return DayStateFuture
	}

	if param.WearHours == 0 && param.SyncCount == 0 {
		return DayStateNoActivity
	}

	if param.WearHours > 0 && param.SyncCount == 0 && param.Date == param.AsOf {
		return DayStateWaitingToSync
	}

	if param.WearHours >= MinWearHours && param.SyncCount > 0 {
		return DayStateMet
	}

	return DayStateUnmet
}

func IsWearHoursMet(wearHours float32) bool {
	return wearHours >= MinWearHours
}

func IsSyncCountMet(syncCount uint) bool {
	return syncCount > 0
}

func FirstSunday(date string) (time.Time, error) {
	parsedDate, err := time.Parse(DateLayout, date)
	if err != nil {
		return time.Time{}, err
	}

	firstDayTime := time.Date(
		parsedDate.Year(), 
		parsedDate.Month(), 
		1, 0, 0, 0, 0, 
		parsedDate.Location(),
	)

	daysToSubtract := int(firstDayTime.Weekday())

	return firstDayTime.AddDate(0, 0, -daysToSubtract), nil
}

func LastSaturday(date string) (time.Time, error) {
	parsedDate, err := time.Parse(DateLayout, date)
	if err != nil {
		return time.Time{}, err
	}

	lastDayTime := time.Date(
		parsedDate.Year(), 
		parsedDate.Month() + 1, 
		0, 23, 59, 59, 999999999,
		parsedDate.Location(),
	)

	daysToAdd := int(time.Saturday) - int(lastDayTime.Weekday())

	return lastDayTime.AddDate(0, 0, daysToAdd), nil
}

func ListDays(date string) ([]time.Time, error) {
	firstSunday, err := FirstSunday(date)
	if err != nil {
		return nil, err
	}

	lastSaturday, err := LastSaturday(date)
	if err != nil {
		return nil, err
	}

	var days []time.Time
	current := firstSunday
	for !current.After(lastSaturday) {
		days = append(days, current)
		current = current.AddDate(0, 0, 1)
	}

	return days, nil
}