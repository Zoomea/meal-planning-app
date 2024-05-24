package db

import "context"

// A generic database interface
// (C)reate, (R)ead, (U)pdate, (D)elete, (L)ist.
type Crudler[T any] interface {
	Create(ctx context.Context, v T) (int64, error)
	// Doesn't return an error if the item isn't found, it just doesn't
	// include it in the output
	Read(ctx context.Context, ids []int64) ([]T, error)
	Update(ctx context.Context, id int64, v T) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]T, error)
}

// Schedule contains pointers to recipes
// The "type" is often "breakfast", "lunch", or "dinner"
type Schedule struct {
	Date    Date
	Type    string
	Recipes []int64
}

type Date struct {
	Day   int64
	Month int64
	Year  int64
}

type ScheduleStore interface {
	List(ctx context.Context, start, end Date) ([]Schedule, error)
	UpsertSchedule(context.Context, Schedule) error
}
