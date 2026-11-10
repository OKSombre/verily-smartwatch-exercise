package builder

import (
	"slices"
	"sort"
	"sync"

	"github.com/verily/calendar/api/internal/utils"
)

type RecordStore struct {
	Mu      sync.Mutex
	Participants []Participant
	Records map[string][]DailyRecord
	AsOfDate string
}

func NewRecordStore() *RecordStore {
	return &RecordStore{
		Participants: make([]Participant, 0),
		Records: make(map[string][]DailyRecord),
		AsOfDate: "2026-10-21",
	}
}

func NewRecordStoreWithSeed(asOfDate string) *RecordStore {
	daily, participants, err := BuildSeedData(asOfDate)
	if err != nil {
		return NewRecordStore()
	}
	
	records := make(map[string][]DailyRecord)
	for _, participant := range participants {
		records[participant.ParticipantID] = make([]DailyRecord, 0)
	}
	for _, record := range daily {
		records[record.ParticipantID] = append(records[record.ParticipantID], record)
	}
	return &RecordStore{
		Participants: participants,
		Records: records,
		AsOfDate: asOfDate,
	}
}

// AllParticipants returns all participants.
func (r *RecordStore) AllParticipants() []Participant {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	return r.Participants
}

// GetParticipant returns a participant by ID.
func (r *RecordStore) GetParticipant(id string) (Participant, bool) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	for _, participant := range r.Participants {
		if participant.ParticipantID == id {
			return participant, true
		}
	}
	return Participant{}, false
}

// AllRecords returns all records.
func (r *RecordStore) AllRecords() []DailyRecord {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	list := make([]DailyRecord, 0)
	for _, records := range r.Records {
		list = append(list, records...)
	}
	sort.Slice(list, func(i, j int) bool { 
		return list[i].ParticipantID < list[j].ParticipantID
	})
	return list
}

// GetRecords returns the records for a participant.
func (r *RecordStore) GetRecords(id string) ([]DailyRecord, bool) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	records, ok := r.Records[id]
	if !ok {
		return nil, false
	}
	return records, true
}

// GetMonth returns the records for a participant in a given month in Sun-Sat grid.
func (r *RecordStore) GetRecordsByMonth(id string, month string) ([]DailyRecord, bool) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	records, ok := r.Records[id]
	if !ok {
		return nil, false
	}
	monthDays, err := utils.ListDays(month)
	if err != nil {
		return nil, false
	}
	monthRecords := make([]DailyRecord, 0)
	for _, record := range records {
		if slices.Contains(monthDays, record.Date) {
			monthRecords = append(monthRecords, record)
		}
	}

	return monthRecords, true
}