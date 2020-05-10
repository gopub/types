package types

import "time"

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Date time.Time

func NowDate() *Date {
	t := time.Now()
	return (*Date)(&t)
}

func NewDate(y, m, d int) *Date {
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	return (*Date)(&t)
}

func (d *Date) Weekday() int {
	return int((*time.Time)(d).Weekday())
}

func (d *Date) Year() int {
	return (*time.Time)(d).Year()
}

func (d *Date) Month() int {
	return int((*time.Time)(d).Month())
}

func (d *Date) Day() int {
	return (*time.Time)(d).Day()
}

func (d *Date) Add(years, months, days int) *Date {
	dd := (time.Time)(*d).AddDate(years, months, days)
	return (*Date)(&dd)
}

func (d *Date) NextDay() *Date {
	return d.Add(0, 0, 1)
}

func (d *Date) LastDay() *Date {
	return d.Add(0, 0, -1)
}

func (d *Date) NextMonth() *Date {
	return d.Add(0, 1, 0)
}

func (d *Date) LastMonth() *Date {
	return d.Add(0, -1, 0)
}

func (d *Date) FirstDayOfMonth() *Date {
	return NewDate(d.Year(), d.Month(), 1)
}

func (d *Date) LastDayOfMonth() *Date {
	return d.NextMonth().FirstDayOfMonth().LastDay()
}

func (d *Date) DaysOfMonth() int {
	return d.LastDayOfMonth().Day() - d.FirstDayOfMonth().Day() + 1
}
