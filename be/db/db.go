package db

import "context"

// A generic database interface
// (C)reate, (R)ead, (U)pdate, (D)elete, (L)ist.
type Crudler[T IDer] interface {
	Create(ctx context.Context, v T) (int64, error)
	Read(ctx context.Context, ids []int64) ([]T, error)
	Update(ctx context.Context, id int64, v T) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]T, error)
}

type IDer interface {
	SetID(int64)
}

type ID int64

func (i *ID) SetID(id int64) {
	*i = ID(id)
}
