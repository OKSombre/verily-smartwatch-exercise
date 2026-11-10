package api

import "fmt"

type DayState int

const (
	DayStateFuture DayState = iota
	DayStateNoActivity
	DayStateWaitingToSync
	DayStateMet
	DayStateUnmet
)

const MinWearHours = 16

type DayStateParam struct {
	Date      string
	AsOf      string
	WearHours float32
	SyncCount uint
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