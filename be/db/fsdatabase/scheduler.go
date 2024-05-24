package fsdatabase

import (
	"context"

	"github.com/Zoomea/meal-planning-app/db"
)

type scheduleDB struct {
	scheds map[dateWithTime]db.Schedule
}

// For instance, this could be 2024-04-02 and "breakfast"
type dateWithTime struct {
	Date db.Date
	Time string
}

// Returns a database handler than implements the ScheduleStore interface
func NewScheduleDB() *scheduleDB {
	return &scheduleDB{make(map[dateWithTime]db.Schedule)}
}

func (s *scheduleDB) List(ctx context.Context, start, end db.Date) ([]db.Schedule, error) {
	resScheds := make([]db.Schedule, 0)

	for _, sched := range s.scheds {
		if inDateRange(start, end, sched.Date) {
			resScheds = append(resScheds, sched)
		}
	}

	return resScheds, nil
}

func (s *scheduleDB) UpsertSchedule(ctx context.Context, sched db.Schedule) error {
	dt := dateWithTime{sched.Date, sched.Type}
	s.scheds[dt] = sched
	return nil
}

// Checks whether "d" is within the date range of start:end (inclusive on both
// ends).
func inDateRange(start, end, d db.Date) bool {
	return start.Year <= d.Year && d.Year <= end.Year &&
		start.Month <= d.Month && d.Month <= end.Month &&
		start.Day <= d.Day && d.Day <= end.Day
}
