package main

import (
	"fmt"

	"github.com/verily/calendar/internal/builder"
)

const AsOfDate = "2026-10-21"

func main() {
    daily, participants, err := builder.BuildSeedData(AsOfDate)
    if err != nil {
        panic(err)
    }
    fmt.Println(len(daily))
	fmt.Println(daily[0])

    fmt.Println(len(participants))
	fmt.Println(participants[0])

	month := AsOfDate[:7]
	participantMonthRecord := builder.BuildParticipantMonthRecord(
		daily,
		month,
		participants[1],
	)
	builder.DaystatesForParticipantMonth(participantMonthRecord, AsOfDate)
}