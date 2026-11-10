package utils

import "time"

type DayState string

const (
	DayStateFuture DayState = "FUTURE"
	DayStateNoActivity DayState = "NO_ACTIVITY"
	DayStateWaitingToSync DayState = "WAITING_TO_SYNC"
	DayStateMet DayState = "MET"
	DayStateUnmet DayState = "UNMET"
)

const DateLayout = "2006-01-02"
const MonthLayout = "2006-01"

type DayStateParam struct {
	Date      string 	// YYYY-MM-DD format
	AsOf      string 	// YYYY-MM-DD format
	WearHours float32 	// Total hours the device was worn
	SyncCount uint 		// Number of times the device was synced
}

// GetDayState returns the state for a given day.
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

// IsWearHoursMet returns whether the wear hours meet the minimum requirement.
func IsWearHoursMet(wearHours float32) bool {
	return wearHours >= MinWearHours
}

// IsSyncCountMet returns whether the sync count meets the minimum requirement.
func IsSyncCountMet(syncCount uint) bool {
	return syncCount >= MinSyncCount
}

// FirstSunday returns the first Sunday of the month.
func FirstSunday(month string) (time.Time, error) {
	parsedDate, err := time.Parse(MonthLayout, month)
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

// LastSaturday returns the last Saturday of the month.
func LastSaturday(month string) (time.Time, error) {
	parsedDate, err := time.Parse(MonthLayout, month)
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

// ListDays returns all dates from the first Sunday to the last Saturday of the month.
func ListDays(month string) ([]string, error) {
	firstSunday, err := FirstSunday(month)
	if err != nil {
		return nil, err
	}

	lastSaturday, err := LastSaturday(month)
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