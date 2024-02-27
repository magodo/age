package age

import (
	"time"
)

type T struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Sec   int
}

func (t T) after(ot T) bool {
	if t.Year != ot.Year {
		return t.Year > ot.Year
	}
	if t.Month != ot.Month {
		return t.Month > ot.Month
	}
	if t.Day != ot.Day {
		return t.Day > ot.Day
	}
	if t.Hour != ot.Hour {
		return t.Hour > ot.Hour
	}
	if t.Min != ot.Min {
		return t.Min > ot.Min
	}
	if t.Sec != ot.Sec {
		return t.Sec > ot.Sec
	}
	return false
}

func Age(birth, target time.Time) int {
	if birth.After(target) {
		return -1
	}
	birthTime := parse(birth)
	targetTime := parse(target)

	res := targetTime.Year - birthTime.Year
	if res == 0 {
		return 0
	}
	targetTime.Year = birthTime.Year
	if birthTime.after(targetTime) {
		res -= 1
	}
	return res
}

func parse(t time.Time) T {
	tt, ok := parseRFC3339(t.UTC().Format(time.RFC3339))
	if !ok {
		panic("invalid time")
	}
	return tt
}
