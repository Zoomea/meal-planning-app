package fsdatabase

import (
	"context"

	"github.com/Zoomea/meal-planning-app/db"
)

type ScheduleDB struct {
	scheds map[db.Date]db.Schedule
}

func NewScheduleDB() ScheduleDB {
	return ScheduleDB{make(map[db.Date]db.Schedule)}
}

func (s *ScheduleDB) List(ctx context.Context, start, end db.Date) ([]db.Schedule, error) {
	resScheds := make([]db.Schedule, 0)

	for _, sched := range s.scheds {
		if inDateRange(start, end, sched.Date) {
			resScheds = append(resScheds, sched)
		}
	}

	return resScheds, nil
}

func (s *ScheduleDB) UpdateSchedule(ctx context.Context, date db.Date, sched db.Schedule) error {
	s.scheds[date] = sched
	return nil
}

// Checks whether "d" is within the date range of start:end (inclusive on both
// ends).
func inDateRange(start, end, d db.Date) bool {
	return start.Year <= d.Year && d.Year <= end.Year &&
		start.Month <= d.Month && d.Month <= end.Month &&
		start.Day <= d.Day && d.Day <= end.Day
}
