package date_v1

import (
	"time"
)

func New(t time.Time) *Date {
	return &Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func Now() *Date {
	return New(time.Now())
}

func (t *Date) AsTime() time.Time {
	y := t.Year
	m := t.Month
	d := t.Day
	if y == 0 {
		y = 1
	}
	if m == 0 {
		m = 1
	}
	if d == 0 {
		d = 1
	}
	return time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.Local)
}
