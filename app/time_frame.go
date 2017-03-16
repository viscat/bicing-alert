package app

import "time"

type TimeFrame struct {
	weekDays []time.Weekday
	dayFrame struct {
		from time.Time
		to time.Time
	}
}