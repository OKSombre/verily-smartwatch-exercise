package api_test

import (
	"testing"

	"github.com/verily/calendar/api"
)

func TestGetDayState(t *testing.T) {
	tests := []struct {
		name     string
		param    api.DayStateParam
		expected api.DayState
	}{
		{
			name:     "Future",
			param:    api.DayStateParam{Date: "2024-06-02", AsOf: "2024-06-01", WearHours: 8.5, SyncCount: 3},
			expected: api.DayStateFuture,
		},
		{
			name:     "FutureNoData",
			param:    api.DayStateParam{Date: "2024-06-02", AsOf: "2024-06-01"},
			expected: api.DayStateFuture,
		},
		{
			name:     "NoActivity",
			param:    api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 0, SyncCount: 0},
			expected: api.DayStateNoActivity,
		},
		{
			name:     "WaitingToSync",
			param:    api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-01", WearHours: 8.5, SyncCount: 0},
			expected: api.DayStateWaitingToSync,
		},
		{
			name:     "Met",
			param:    api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: api.MinWearHours, SyncCount: 1},
			expected: api.DayStateMet,
		},
		{
			name:     "UnmetNoSync",
			param:    api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: api.MinWearHours, SyncCount: 0},
			expected: api.DayStateUnmet,
		},
		{
			name:     "UnmetNoWear",
			param:    api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 0, SyncCount: 1},
			expected: api.DayStateUnmet,
		},
		{
			name:     "UnmetNotMinimumWear",
			param:    api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: api.MinWearHours - 0.1, SyncCount: 1},
			expected: api.DayStateUnmet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := api.GetDayState(tt.param)
			if got != tt.expected {
				t.Errorf("GetDayState() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsWearHoursMet(t *testing.T) {
	tests := []struct {
		name      string
		wearHours float32
		expected  bool
	}{
		{
			name:      "Met",
			wearHours: api.MinWearHours,
			expected:  true,
		},
		{
			name:      "NotMet",
			wearHours: api.MinWearHours - 0.1,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := api.IsWearHoursMet(tt.wearHours)
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
			got := api.IsSyncCountMet(tt.syncCount)
			if got != tt.expected {
				t.Errorf("IsSyncCountMet() = %v, want %v", got, tt.expected)
			}
		})
	}
}