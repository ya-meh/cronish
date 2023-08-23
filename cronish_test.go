package cronish

import (
	"math/rand"
	"testing"
	"time"
)

func TestT_Time(test *testing.T) {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 100000; i++ {
		timeFormat := LayoutsTime[rand.Int63n(int64(len(LayoutsTime)))]
		dayFormat := LayoutsDay[rand.Int63n(int64(len(LayoutsDay)))]

		t := time.Unix(rand.Int63n(delta)+min, 0)
		closest := Parse(DayTry(t.Format(dayFormat)), TimeTry(t.Format(timeFormat)))

		if closest.Before(time.Now()) {
			test.Fatalf("closest.Before(time.Now()), %s", t.Format(time.DateTime))
		}

		if closest.Weekday() != t.Weekday() || closest.Format(timeFormat) != t.Format(timeFormat) {
			test.Logf("Time: %s, Day: %s", timeFormat, dayFormat)
			test.Logf("day-time broken, restr: %s, got: %s\n%s =? %s", t.Format(time.DateTime), closest.Format(time.DateTime), t.Weekday().String(), closest.Weekday().String())
			_ = Parse(Day(dayFormat, t.Format(dayFormat)), TimeTry(t.Format(timeFormat)))
			test.Fail()
		}
	}

	for i := 0; i < 100000; i++ {
		timeFormat := LayoutsTime[rand.Int63n(int64(len(LayoutsTime)))]
		dateFormat := LayoutsDate[rand.Int63n(int64(len(LayoutsDate)))]

		t := time.Unix(rand.Int63n(delta)+min, 0)
		closest := Parse(DateTry(t.Format(dateFormat)), TimeTry(t.Format(timeFormat)))

		if closest.Before(time.Now()) {
			test.Fatalf("closest.Before(time.Now()), %s", t.Format(time.DateTime))
		}

		if closest.Day() != t.Day() || closest.Month() != t.Month() || closest.Format(timeFormat) != t.Format(timeFormat) {
			test.Logf("Time: %s, Date: %s", timeFormat, dateFormat)
			test.Logf("Time: %s, Date: %s", t.Format(timeFormat), t.Format(dateFormat))
			test.Logf("Time: %s, Date: %s", closest.Format(timeFormat), closest.Format(dateFormat))
			_ = Parse(Date(dateFormat, t.Format(dateFormat)), TimeTry(t.Format(timeFormat)))
			test.Fatalf("date-time broken, restr: %s, got: %s", t.Format(time.DateTime), closest.Format(time.DateTime))
		}
	}
}

func BenchmarkT_Time(b *testing.B) {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < b.N; i++ {
		t := time.Unix(rand.Int63n(delta)+min, 0)

		_ = Parse(TimeTry(t.Format(time.TimeOnly)), DateTry(t.Format(time.DateOnly)))
	}
}
