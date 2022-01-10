package util

import (
	"time"

	"google.golang.org/api/classroom/v1"
)

func ParseClassroomTime(d *classroom.Date, t *classroom.TimeOfDay) *time.Time {
	if d == nil || t == nil {
		// not due
		return nil
	}

	out := time.Date(
		int(d.Year), time.Month(d.Month), int(d.Day),
		int(t.Hours), int(t.Minutes), int(t.Seconds), int(t.Nanos),
		time.UTC,
	)

	return &out
}
