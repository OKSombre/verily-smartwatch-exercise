package main

import (
	"fmt"

	"github.com/verily/calendar/api"
)

func main() {
    api.Hello()
    api.GetDayState(api.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 8.5, SyncCount: 3})
	fmt.Println(api.FirstSunday("2026-11"))
	fmt.Println(api.LastSaturday("2026-11"))
}