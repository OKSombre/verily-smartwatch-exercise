package main

import (
	"fmt"

	"github.com/verily/calendar/internal/utils"
)

func main() {
    utils.Hello()
    utils.GetDayState(utils.DayStateParam{Date: "2024-06-01", AsOf: "2024-06-02", WearHours: 8.5, SyncCount: 3})
	fmt.Println(utils.FirstSunday("2026-11"))
	fmt.Println(utils.LastSaturday("2026-11"))
}