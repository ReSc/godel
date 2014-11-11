package main

import (
	"github.com/ReSc/fmt"
	"github.com/ReSc/godel/core/cron"
)
var  days := cron.WeekDays

func main() {
 	name := days.Sunday.String()
	day, _ := cron.ParseWeekDays(name)
	if day == days.Sunday {
		fmt.Printline("success")
	} else {
		fmt.Printline("error")
	}
}
