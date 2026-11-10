package builder

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/verily/calendar/internal/utils"
)

//go:embed seed/daily.json
//go:embed seed/participants.json
var seed embed.FS

type Wear struct {
	Hours float32 `json:"wearHours"`
	Met bool
}

type Sync struct {
	Count uint `json:"syncCount"`
	Met bool
}

type DailyRecord struct {
	ParticipantID string `json:"participantId"`
    Date string `json:"date"`
    WearHours Wear
    SyncCount Sync
}

type DailyRecordJson struct {
	ParticipantID string `json:"participantId"`
    Date string `json:"date"`
    WearHours float32 `json:"wearHours"`
    SyncCount uint `json:"syncCount"`
}

type Participant struct {
	ParticipantID string `json:"participantId"`
	DeviceID string `json:"deviceId"`
	Site string `json:"site"`
}

type ParticipantMonthRecord struct {
	ParticipantID string
	DeviceID string
	Site string
	Records map[string]DailyRecord
	Month string
}

func ParseDaily(data []byte, asOfDate string) ([]DailyRecord, error) {
	var records []DailyRecordJson
	err := json.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}

	dailyRecords := make([]DailyRecord, len(records))
	for i, record := range records {
		dailyRecords[i] = DailyRecord{
			ParticipantID: record.ParticipantID,
			Date:          record.Date,
			WearHours: Wear{
				Hours: record.WearHours,
				Met:   utils.IsWearHoursMet(record.WearHours),
			},
			SyncCount: Sync{
				Count: record.SyncCount,
				Met:   utils.IsSyncCountMet(record.SyncCount),
			},
		}

		if (utils.GetDayState(utils.DayStateParam{
			Date:      record.Date,
			AsOf:      asOfDate,
			WearHours: record.WearHours,
			SyncCount: record.SyncCount,
		}) == utils.DayStateFuture) {
			dailyRecords[i].WearHours.Hours = 0
			dailyRecords[i].WearHours.Met = false

			dailyRecords[i].SyncCount.Count = 0
			dailyRecords[i].SyncCount.Met = false
		}
	}
	
	return dailyRecords, err
}

func ParseParticipants(data []byte) ([]Participant, error) {
	var participants []Participant
	err := json.Unmarshal(data, &participants)
	return participants, err
}

func BuildSeedData(asOfDate string) ([]DailyRecord, []Participant, error) {
	dailyData, err := seed.ReadFile("seed/daily.json")
	if err != nil {
		return nil, nil, err
	}

	daily, err := ParseDaily(dailyData, asOfDate)
	if err != nil {
		return nil, nil, err
	}

	participantData, err := seed.ReadFile("seed/participants.json")
	if err != nil {
		return nil, nil, err
	}

	participants, err := ParseParticipants(participantData)
	if err != nil {
		return nil, nil, err
	}

	return daily, participants, nil
}

func BuildParticipantMonthRecord(
	daily []DailyRecord,  
	month string, 
	participant Participant,
	) ParticipantMonthRecord {
	monthDays, err := utils.ListDays(month)
	if err != nil {
		return ParticipantMonthRecord{}
	}

	recordsByDate := make(map[string]DailyRecord, len(monthDays))
	for _, date := range monthDays {
		recordsByDate[date] = DailyRecord{
			ParticipantID: participant.ParticipantID,
			Date:          date,
			WearHours:     Wear{},
			SyncCount:     Sync{},
		}
	}
	for _, record := range daily {
		if record.ParticipantID == participant.ParticipantID {
			if _, ok := recordsByDate[record.Date]; ok {
				r := recordsByDate[record.Date]
				r.WearHours.Hours += record.WearHours.Hours
				r.WearHours.Met = utils.IsWearHoursMet(r.WearHours.Hours)
				r.SyncCount.Count += record.SyncCount.Count
				r.SyncCount.Met = utils.IsSyncCountMet(r.SyncCount.Count)
				recordsByDate[record.Date] = r
			}
		}
	}

	return ParticipantMonthRecord{
		ParticipantID: participant.ParticipantID,
		DeviceID:      participant.DeviceID,
		Site:          participant.Site,
		Records:       recordsByDate,
		Month:         month,
	}
}

func DaystatesForParticipantMonth(record ParticipantMonthRecord, asOfDate string) string {
	dayStates := make(map[utils.DayState]int)
	for date, dailyRecord := range record.Records {
		if _, ok := dayStates[utils.GetDayState(utils.DayStateParam{
			Date:      date,
			AsOf:      asOfDate,
			WearHours: dailyRecord.WearHours.Hours,
			SyncCount: dailyRecord.SyncCount.Count,
		})]; !ok {
			dayStates[utils.GetDayState(utils.DayStateParam{
				Date:      date,
				AsOf:      asOfDate,
				WearHours: dailyRecord.WearHours.Hours,
				SyncCount: dailyRecord.SyncCount.Count,
			})] = 0
		}
		dayStates[utils.GetDayState(utils.DayStateParam{
			Date:      date,
			AsOf:      asOfDate,
			WearHours: dailyRecord.WearHours.Hours,
			SyncCount: dailyRecord.SyncCount.Count,
		})]++
	}

	return fmt.Sprintf("MET: %d, UNMET: %d, NO ACTIVITY: %d, WAITING TO SYNC: %d, FUTURE: %d",
		dayStates[utils.DayStateMet],
		dayStates[utils.DayStateUnmet],
		dayStates[utils.DayStateNoActivity],
		dayStates[utils.DayStateWaitingToSync],
		dayStates[utils.DayStateFuture],
	)
}