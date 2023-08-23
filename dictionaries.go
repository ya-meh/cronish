package cronish

import (
	"strings"
	"time"
)

type Dictionary map[string]time.Weekday

func (d Dictionary) And(other Dictionary) Dictionary {
	result := Dictionary{}
	for k, v := range d {
		result[k] = v
	}
	for k, v := range other {
		result[k] = v
	}
	return result
}

func (d Dictionary) Get(name string) time.Weekday {
	day, _ := d.TryGet(name)
	return day
}

func (d Dictionary) TryGet(name string) (time.Weekday, bool) {
	if day, ok := d[strings.ToLower(name)]; ok {
		return day, true
	}

	return time.Sunday, false
}

var (
	Russian = Dictionary{
		"понедельник": time.Monday,
		"вторник":     time.Tuesday,
		"среда":       time.Wednesday,
		"четверг":     time.Thursday,
		"пятница":     time.Friday,
		"суббота":     time.Saturday,
		"воскресенье": time.Sunday,
		"пн":          time.Monday,
		"вт":          time.Tuesday,
		"ср":          time.Wednesday,
		"чт":          time.Thursday,
		"пт":          time.Friday,
		"сб":          time.Saturday,
		"вс":          time.Sunday,
	}

	English = Dictionary{
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
		"sunday":    time.Sunday,
		"mon":       time.Monday,
		"tue":       time.Tuesday,
		"wed":       time.Wednesday,
		"thu":       time.Thursday,
		"fri":       time.Friday,
		"sat":       time.Saturday,
		"sun":       time.Sunday,
	}

	Dictionaries = []Dictionary{English, Russian}
)
