package utils

import "time"

type DayState int

const (
	DayStateFuture DayState = iota
	DayStateNoActivity
	DayStateWaitingToSync
	DayStateMet
	DayStateUnmet
)

const DateLayout = "2006-01-02"
const MonthLayout = "2006-01"

type DayStateParam struct {
	Date      string 	// YYYY-MM-DD format
	AsOf      string 	// YYYY-MM-DD format
	WearHours float32 	// Total hours the device was worn
	SyncCount uint 		// Number of times the device was synced
}

func GetDayState(param DayStateParam) DayState {
	if param.Date > param.AsOf {
		return DayStateFuture
	}

	if param.WearHours == 0 && param.SyncCount == 0 {
		return DayStateNoActivity
	}

	if param.WearHours > 0 && param.SyncCount == 0 && param.Date == param.AsOf {
		return DayStateWaitingToSync
	}

	if IsWearHoursMet(param.WearHours) && IsSyncCountMet(param.SyncCount) {
		return DayStateMet
	}

	return DayStateUnmet
}

func IsWearHoursMet(wearHours float32) bool {
	return wearHours >= MinWearHours
}

func IsSyncCountMet(syncCount uint) bool {
	return syncCount >= MinSyncCount
}

func FirstSunday(date string) (time.Time, error) {
	parsedDate, err := time.Parse(MonthLayout, date)
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
	parsedDate, err := time.Parse(MonthLayout, date)
	if err != nil {
		return time.Time{}, err
	}

	lastDayTime := time.Date(
		parsedDate.Year(), 
		parsedDate.Month() + 1, 
		0, 0, 0, 0, 0,
		parsedDate.Location(),
	)

	daysToAdd := int(time.Saturday) - int(lastDayTime.Weekday())

	return lastDayTime.AddDate(0, 0, daysToAdd), nil
}

func ListDays(date string) ([]string, error) {
	firstSunday, err := FirstSunday(date)
	if err != nil {
		return nil, err
	}

	lastSaturday, err := LastSaturday(date)
	if err != nil {
		return nil, err
	}

	var days []string
	current := firstSunday
	for !current.After(lastSaturday) {
		days = append(days, current.Format(DateLayout))
		current = current.AddDate(0, 0, 1)
	}

	return days, nil
}