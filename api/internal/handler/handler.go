package handler

import (
	"net/http"

	"github.com/verily/calendar/api/internal/builder"
	"github.com/verily/calendar/api/internal/utils"
)

type Handler struct {
	Store *builder.RecordStore
}

type WearJSON struct {
	Hours 	float32 `json:"hours"`
	Met 	bool 	`json:"met"`
}

type SyncJSON struct {
	Count 	uint 	`json:"count"`
	Met 	bool 	`json:"met"`
}

type DailyRecordJSON struct {
	Date      string 	`json:"date"`	// YYYY-MM-DD format
	State     string 	`json:"state"`	// One of FUTURE, NO_ACTIVITY, WAITING_TO_SYNC, MET, UNMET
	Wear      WearJSON 	`json:"wear"`	// Wear data for the day
	Sync      SyncJSON 	`json:"sync"`	// Sync data for the day
}

type MonthlyRecordJSON struct {
	ParticipantID 	string				`json:"participantId"`	// Participant ID
	Site 			string				`json:"site"` 			// Site name
	Month  			string              `json:"month"`			// YYYY-MM format
	Days 			[]DailyRecordJSON 	`json:"days"`			// Daily records for the month
}

// GET /api/studies/alpha/participants/{participantID}/calendar?month=YYYY-MM
// Returns the calendar records for a specific participant and month.
func (h *Handler) GetParticipantRecordsByMonth(w http.ResponseWriter, r *http.Request) {
	participantID, err := pathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	month, err := queryMonth(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	monthlyRecords, ok := h.Store.GetRecordsByMonth(participantID, month)
	if !ok {
		writeError(w, http.StatusNotFound, "participant not found")
		return
	}

	records := h.CreateMonthlyRecordJSON(participantID, month, monthlyRecords)
	writeJSON(w, http.StatusOK, records)
}

func (h *Handler) CreateMonthlyRecordJSON(participantID string, month string, days []builder.DailyRecord) MonthlyRecordJSON {
	participant, ok := h.Store.GetParticipant(participantID)
	if !ok {
		return MonthlyRecordJSON{}
	}

	dayList, err := utils.ListDays(month)
	if err != nil {
		return MonthlyRecordJSON{}
	}

	recordsByDate := make(map[string]DailyRecordJSON, len(dayList))
	for _, date := range dayList {
		recordsByDate[date] = DailyRecordJSON{
			Date:  date,
			State: string(utils.GetDayState(utils.DayStateParam{
				Date: date,
				AsOf: h.Store.AsOfDate,
			})),
		}
	}

	for _, record := range days {
		recordsByDate[record.Date] = DailyRecordJSON{
			Date:  record.Date,
			State: string(record.DayState),
			Wear: WearJSON{
				Hours: record.WearHours.Hours,
				Met:   record.WearHours.Met,
			},
			Sync: SyncJSON{
				Count: record.SyncCount.Count,
				Met:   record.SyncCount.Met,
			},
		}
	}

	daysJSON := make([]DailyRecordJSON, 0, len(dayList))
	for _, date := range dayList {
		record := recordsByDate[date]
		daysJSON = append(daysJSON, record)
	}

	return MonthlyRecordJSON{
		ParticipantID: participantID,
		Site:          participant.Site,
		Month:         month,
		Days:          daysJSON,
	}
}