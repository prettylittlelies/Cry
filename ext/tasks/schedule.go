package tasks

import (
	"slices"
	"time"
)

type TimeOfDay struct {
	Hour     int
	Minute   int
	Second   int
	Location *time.Location
}

func (t TimeOfDay) on(day time.Time) time.Time {
	loc := t.Location
	if loc == nil {
		loc = time.UTC
	}
	d := day.In(loc)
	return time.Date(d.Year(), d.Month(), d.Day(), t.Hour, t.Minute, t.Second, 0, loc)
}

type schedule interface {
	next(after time.Time) time.Time
}

type interval time.Duration

func (i interval) next(after time.Time) time.Time {
	return after.Add(time.Duration(i))
}

type daily []TimeOfDay

func (d daily) next(after time.Time) time.Time {
	var candidates []time.Time
	for _, t := range d {
		at := t.on(after)
		if !at.After(after) {
			at = t.on(after.AddDate(0, 0, 1))
		}
		candidates = append(candidates, at)
	}
	return slices.MinFunc(candidates, func(a, b time.Time) int { return a.Compare(b) })
}
