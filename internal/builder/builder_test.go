package builder_test

import (
	"testing"

	"github.com/verily/calendar/internal/builder"
)

const AsOfDate = "2026-10-21"

func TestBuildSeedData(t *testing.T) {
	daily, participants, err := builder.BuildSeedData(AsOfDate)
	if err != nil {
		t.Fatalf("BuildSeedData() error = %v", err)
	}
	if len(daily) != 313 {
		t.Errorf("BuildSeedData() returned %d daily records, want 313", len(daily))
	}
	if len(participants) != 14 {
		t.Errorf("BuildSeedData() returned %d participants, want 14", len(participants))
	}
}

func TestBuildParticipantMonthRecord(t *testing.T) {
	daily, participants, err := builder.BuildSeedData(AsOfDate)
	if err != nil {
		t.Fatalf("BuildSeedData() error = %v", err)
	}

	tests := []struct {
		participant builder.Participant
		month       string
		expect      string
	}{
		{
			participant: participants[0], // A1000
			month:       "2026-10",
			expect: 	"MET: 0, UNMET: 0, NO ACTIVITY: 25, WAITING TO SYNC: 0, FUTURE: 10",
		},
		{
			participant: participants[1], // A1001
			month:       "2026-10",
			expect: 	"MET: 17, UNMET: 4, NO ACTIVITY: 3, WAITING TO SYNC: 1, FUTURE: 10",
		},
		{
			participant: participants[1], // A1001
			month:       "2026-09",
			expect: 	"MET: 5, UNMET: 1, NO ACTIVITY: 29, WAITING TO SYNC: 0, FUTURE: 0",
		},
		{
			participant: participants[9], // A1009
			month:       "2026-10",
			expect: 	"MET: 0, UNMET: 25, NO ACTIVITY: 0, WAITING TO SYNC: 0, FUTURE: 10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.participant.ParticipantID+"_"+tt.month, func(t *testing.T) {
			record := builder.BuildParticipantMonthRecord(
				daily, 
				tt.month,
				tt.participant, 
			)
			got := builder.DaystatesForParticipantMonth(record, AsOfDate)
			if got != tt.expect {
				t.Errorf("DaystatesForParticipantMonth() = %q, want %q", got, tt.expect)
			}

		})
	}
}