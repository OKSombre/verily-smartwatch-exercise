package utils_test

import (
	"testing"
	"time"

	"github.com/verily/calendar/api/internal/utils"
)

func TestGetDayState(t *testing.T) {
	tests := []struct {
		name     string
		param    utils.DayStateParam
		expected utils.DayState
	}{
		{
			name:     "Future",
			param:    utils.DayStateParam{Date: "2024-06-02", AsOf: "2024-06-01", WearHours: 8.5, SyncCount: 3},
			expected: utils.DayStateFuture,
		},
		{
			name:     "FutureNoData",
			param:    utils.DayStateParam{Date: "2024-06-02", AsOf: "2024-06-01"},
			expected: utils.DayStateFuture,
		},
		{
			name:     "NoActivity",
			param:    utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 0, SyncCount: 0},
			expected: utils.DayStateNoActivity,
		},
		{
			name:     "WaitingToSync",
			param:    utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-01", WearHours: 8.5, SyncCount: 0},
			expected: utils.DayStateWaitingToSync,
		},
		{
			name:     "Met",
			param:    utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 16, SyncCount: 1},
			expected: utils.DayStateMet,
		},
		{
			name:     "UnmetNoSync",
			param:    utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 16, SyncCount: 0},
			expected: utils.DayStateUnmet,
		},
		{
			name:     "UnmetNoWear",
			param:    utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 0, SyncCount: 1},
			expected: utils.DayStateUnmet,
		},
		{
			name:     "UnmetNotMinimumWear",
			param:    utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 15.9, SyncCount: 1},
			expected: utils.DayStateUnmet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GetDayState(tt.param)
			if got != tt.expected {
				t.Errorf("GetDayState() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsWearHoursMet(t *testing.T) {
	tests := []struct {
		name      	string
		wearHours   float32
		expected  	bool
	}{
		{
			name:      "Met",
			wearHours:    16,
			expected:  true,
		},
		{
			name:      "NotMet",
			wearHours:    15.9,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.IsWearHoursMet(tt.wearHours)
			if got != tt.expected {
				t.Errorf("IsWearHoursMet() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsSyncCountMet(t *testing.T) {
	tests := []struct {
		name      string
		syncCount uint
		expected  bool
	}{
		{
			name:      "Met",
			syncCount: 1,
			expected:  true,
		},
		{
			name:      "NotMet",
			syncCount: 0,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.IsSyncCountMet(tt.syncCount)
			if got != tt.expected {
				t.Errorf("IsSyncCountMet() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalendarHelpers(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expectedFirstSunday time.Time
		expectedLastSaturday time.Time
		expectedDays int
	}{
		{
			name:     "February 2024 (Leap Year)",
			date:     "2024-02",
			expectedFirstSunday: time.Date(2024, time.January, 28, 0, 0, 0, 0, time.UTC),
			expectedLastSaturday: time.Date(2024, time.March, 2, 0, 0, 0, 0, time.UTC),
			expectedDays: 35,
		},
		{
			name:     "February 2026",
			date:     "2026-02",
			expectedFirstSunday: time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC),
			expectedLastSaturday: time.Date(2026, time.February, 28, 0, 0, 0, 0, time.UTC),
			expectedDays: 28,
		},
		{
			name:     "August 2026",
			date:     "2026-08",
			expectedFirstSunday: time.Date(2026, time.July, 26, 0, 0, 0, 0, time.UTC),
			expectedLastSaturday: time.Date(2026, time.September, 5, 0, 0, 0, 0, time.UTC),
			expectedDays: 42,
		},
		{
			name:     "October 2026",
			date:     "2026-10",
			expectedFirstSunday: time.Date(2026, time.September, 27, 0, 0, 0, 0, time.UTC),
			expectedLastSaturday: time.Date(2026, time.October, 31, 0, 0, 0, 0, time.UTC),
			expectedDays: 35,
		},
		{
			name:     "November 2026",
			date:     "2026-11",
			expectedFirstSunday: time.Date(2026, time.November, 1, 0, 0, 0, 0, time.UTC),
			expectedLastSaturday: time.Date(2026, time.December, 5, 0, 0, 0, 0, time.UTC),
			expectedDays: 35,
		},
		{
			name:     "December 2026",
			date:     "2026-12",
			expectedFirstSunday: time.Date(2026, time.November, 29, 0, 0, 0, 0, time.UTC),
			expectedLastSaturday: time.Date(2027, time.January, 2, 0, 0, 0, 0, time.UTC),
			expectedDays: 35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstSunday, err := utils.FirstSunday(tt.date)
			if err != nil {
				t.Fatalf("FirstSunday() error = %v", err)
			}
			if !gotFirstSunday.Equal(tt.expectedFirstSunday) {
				t.Errorf("FirstSunday() = %v, want %v", gotFirstSunday, tt.expectedFirstSunday)
			}

			gotLastSaturday, err := utils.LastSaturday(tt.date)
			if err != nil {
				t.Fatalf("LastSaturday() error = %v", err)
			}
			if !gotLastSaturday.Equal(tt.expectedLastSaturday) {
				t.Errorf("LastSaturday() = %v, want %v", gotLastSaturday, tt.expectedLastSaturday)
			}

			gotDays, err := utils.ListDays(tt.date)
			if err != nil {
				t.Fatalf("ListDays() error = %v", err)
			}
			if len(gotDays) != tt.expectedDays {
				t.Errorf("ListDays() returned %d days, want %d", len(gotDays), tt.expectedDays)
			}
		})
	}
}