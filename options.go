package cronish

import (
	"reflect"
	"strings"
	"time"
	"unicode"
)

var (
	LayoutsTime = []string{"15:04:05", "3:04:05", "15:04", "3:04"}
	LayoutsDate = []string{"02.01.2006", "2.01.2006", "02.01.06", "02.01", "2.01"}
	LayoutsDay  = []string{time.UnixDate, time.RFC850, time.RFC1123Z, time.RFC1123, time.RFC850}
)

type Setter func(*T)

func null(*T) {}

var nullptr = reflect.ValueOf(null).Pointer()

func isNull(s Setter) bool {
	return reflect.ValueOf(s).Pointer() == nullptr
}

func Day(layout, value string, dictionaries ...Dictionary) (setter Setter) {
	begin := strings.Index(layout, "Mon")
	if begin == -1 {
		return null
	}

	end := strings.IndexFunc(value[begin:], func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	if end == -1 {
		end = len(value)
	}

	return DayLiteral(value[begin:end], dictionaries...)
}

func DayLiteral(value string, dictionaries ...Dictionary) (setter Setter) {
	if len(dictionaries) == 0 {
		dictionaries = Dictionaries
	}

	for _, d := range dictionaries {
		day, ok := d.TryGet(value)
		if ok {
			return func(t *T) {
				t.Day = &day
			}
		}
	}

	return null
}

func DayTry(value string, dictionaries ...Dictionary) (setter Setter) {
	for _, layout := range LayoutsDay {
		setter = Day(layout, value, dictionaries...)
		if !isNull(setter) {
			return
		}
	}

	return null
}

func Time(layout, value string) (setter Setter) {
	mod, err := time.Parse(layout, value)
	if err != nil {
		return null
	}

	return func(t *T) {
		t.TimeOpt = new(OptionTime)
		t.TimeOpt.Hour = mod.Hour()
		t.TimeOpt.Minute = mod.Minute()
		t.TimeOpt.Second = mod.Second()
	}
}

func TimeTry(value string) (setter Setter) {
	for _, layout := range LayoutsTime {
		mod, err := time.Parse(layout, value)
		if err == nil {
			return func(t *T) {
				t.TimeOpt = new(OptionTime)
				t.TimeOpt.Hour = mod.Hour()
				t.TimeOpt.Minute = mod.Minute()
				t.TimeOpt.Second = mod.Second()
			}
		}
	}
	return null
}

func Date(layout, value string) (setter Setter) {
	mod, err := time.Parse(layout, value)
	if err != nil {
		return null
	}

	return func(t *T) {
		t.DateOpt = new(OptionDate)
		t.DateOpt.Day = mod.Day()
		t.DateOpt.Month = mod.Month()
		t.DateOpt.Year = mod.Year()
	}
}

func DateTry(value string) (setter Setter) {
	for _, layout := range LayoutsDate {
		mod, err := time.Parse(layout, value)
		if err == nil {
			return func(t *T) {
				t.DateOpt = new(OptionDate)
				t.DateOpt.Day = mod.Day()
				t.DateOpt.Month = mod.Month()
				t.DateOpt.Year = mod.Year()
			}
		}
	}
	return null
}

func All(layout, value string) (setter Setter) {
	_, err := time.Parse(layout, value)
	if err != nil {
		return setter
	}
	day := Day(layout, value)
	time := Time(layout, value)
	date := Date(layout, value)

	return func(t *T) {
		day(t)
		time(t)
		date(t)
	}
}
