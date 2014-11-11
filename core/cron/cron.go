package cron

import (
	strings "strings"
)

// enumMonths is the enum member type of Months
type enumMonths int

// Months is an enumeration
var Months = struct {
	April     enumMonths
	August    enumMonths
	December  enumMonths
	February  enumMonths
	Januari   enumMonths
	July      enumMonths
	June      enumMonths
	March     enumMonths
	May       enumMonths
	November  enumMonths
	October   enumMonths
	September enumMonths
}{
	April:     3,
	August:    7,
	December:  11,
	February:  1,
	Januari:   0,
	July:      6,
	June:      5,
	March:     2,
	May:       4,
	November:  10,
	October:   9,
	September: 8,
}

// String returns the name of the Months enum member
func (this enumMonths) String() string {
	switch this {
	case Months.April:
		return "April"
	case Months.August:
		return "August"
	case Months.December:
		return "December"
	case Months.February:
		return "February"
	case Months.Januari:
		return "Januari"
	case Months.July:
		return "July"
	case Months.June:
		return "June"
	case Months.March:
		return "March"
	case Months.May:
		return "May"
	case Months.November:
		return "November"
	case Months.October:
		return "October"
	case Months.September:
		return "September"
	default:
		panic("Invalid enum member for Months")
	}
}

// ParseMonths returns the value for the name
func ParseMonths(name string) (enumMonths, bool) {
	switch name {
	case "April":
		return Months.April, true
	case "August":
		return Months.August, true
	case "December":
		return Months.December, true
	case "February":
		return Months.February, true
	case "Januari":
		return Months.Januari, true
	case "July":
		return Months.July, true
	case "June":
		return Months.June, true
	case "March":
		return Months.March, true
	case "May":
		return Months.May, true
	case "November":
		return Months.November, true
	case "October":
		return Months.October, true
	case "September":
		return Months.September, true
	default:
		return Months.April, false
	}
}

// ParseMonthsCI returns the value for the name (case insensitive)
func ParseMonthsCI(name string) (enumMonths, bool) {
	name = strings.ToLower(name)
	switch name {
	case "april":
		return Months.April, true
	case "august":
		return Months.August, true
	case "december":
		return Months.December, true
	case "february":
		return Months.February, true
	case "januari":
		return Months.Januari, true
	case "july":
		return Months.July, true
	case "june":
		return Months.June, true
	case "march":
		return Months.March, true
	case "may":
		return Months.May, true
	case "november":
		return Months.November, true
	case "october":
		return Months.October, true
	case "september":
		return Months.September, true
	default:
		return Months.April, false
	}
}

// enumWeekDays is the enum member type of WeekDays
type enumWeekDays int

// WeekDays is an enumeration
var WeekDays = struct {
	Friday    enumWeekDays
	Monday    enumWeekDays
	Saturday  enumWeekDays
	Sunday    enumWeekDays
	Thursday  enumWeekDays
	Tuesday   enumWeekDays
	Wednesday enumWeekDays
}{
	Friday:    5,
	Monday:    1,
	Saturday:  6,
	Sunday:    0,
	Thursday:  4,
	Tuesday:   2,
	Wednesday: 3,
}

// String returns the name of the WeekDays enum member
func (this enumWeekDays) String() string {
	switch this {
	case WeekDays.Friday:
		return "Friday"
	case WeekDays.Monday:
		return "Monday"
	case WeekDays.Saturday:
		return "Saturday"
	case WeekDays.Sunday:
		return "Sunday"
	case WeekDays.Thursday:
		return "Thursday"
	case WeekDays.Tuesday:
		return "Tuesday"
	case WeekDays.Wednesday:
		return "Wednesday"
	default:
		panic("Invalid enum member for WeekDays")
	}
}

// ParseWeekDays returns the value for the name
func ParseWeekDays(name string) (enumWeekDays, bool) {
	switch name {
	case "Friday":
		return WeekDays.Friday, true
	case "Monday":
		return WeekDays.Monday, true
	case "Saturday":
		return WeekDays.Saturday, true
	case "Sunday":
		return WeekDays.Sunday, true
	case "Thursday":
		return WeekDays.Thursday, true
	case "Tuesday":
		return WeekDays.Tuesday, true
	case "Wednesday":
		return WeekDays.Wednesday, true
	default:
		return WeekDays.Friday, false
	}
}

// ParseWeekDaysCI returns the value for the name (case insensitive)
func ParseWeekDaysCI(name string) (enumWeekDays, bool) {
	name = strings.ToLower(name)
	switch name {
	case "friday":
		return WeekDays.Friday, true
	case "monday":
		return WeekDays.Monday, true
	case "saturday":
		return WeekDays.Saturday, true
	case "sunday":
		return WeekDays.Sunday, true
	case "thursday":
		return WeekDays.Thursday, true
	case "tuesday":
		return WeekDays.Tuesday, true
	case "wednesday":
		return WeekDays.Wednesday, true
	default:
		return WeekDays.Friday, false
	}
}
