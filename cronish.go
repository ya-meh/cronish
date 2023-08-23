package cronish

import (
	"fmt"
	"time"
)

type OptionDate struct {
	Day   int
	Month time.Month
	Year  int
}

type OptionTime struct {
	Hour   int
	Minute int
	Second int
}

type T struct {
	Day     *time.Weekday
	DateOpt *OptionDate
	TimeOpt *OptionTime
}

func (t *T) Time() time.Time {
	now := time.Now()

	target := OptionTime{0, 0, 0}
	if t.TimeOpt != nil {
		target = *t.TimeOpt
	}
	diff := (target.Hour-now.Hour())*60*60 + (target.Minute-now.Minute())*60 + (target.Second - now.Second())
	if diff < 0 {
		diff += 24 * 60 * 60
	}
	now = now.Add(time.Duration(diff) * time.Second)

	if t.DateOpt != nil {
		year := t.DateOpt.Year
		if year < now.Year() {
			year = now.Year()
		}

		for {
			if now.Day() == t.DateOpt.Day && now.Month() == t.DateOpt.Month {
				break
			}
			now = now.Add(24 * time.Hour)
		}
	} else if t.Day != nil {
		for {
			if now.Weekday() == *t.Day {
				break
			}
			now = now.Add(24 * time.Hour)
		}
	}

	return now
}

func New(options ...Setter) *T {
	t := T{}
	for _, opt := range options {
		opt(&t)
	}
	return &t
}

func NewSafe(options ...Setter) (*T, error) {
	t := T{}
	for i, opt := range options {
		if isNull(opt) {
			return nil, fmt.Errorf("broken setter, index %d", i)
		}
		opt(&t)
	}
	return &t, nil
}

func Parse(options ...Setter) time.Time {
	return New(options...).Time()
}

func ParseSafe(options ...Setter) (time.Time, error) {
	t, e := NewSafe(options...)
	if e != nil {
		return time.Time{}, e
	}
	return t.Time(), nil
}
