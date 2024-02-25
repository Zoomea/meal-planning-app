// fsdatabase contains a filesystem based database that implements the Crudler
// interface in "db" module.
package fsdatabase

import (
	"context"
	"encoding/json"
	"errors"
	"io"
)

// Implements the Crudler interface in the "db" module.
type repo[T any] struct {
	// The largest ID in the list
	ID    int64       `json:"id"`
	Items map[int64]T `json:"items"`
}

func New[T any]() repo[T] {
	return repo[T]{
		ID:    0,
		Items: make(map[int64]T),
	}
}

func FromJSON[T any](in io.Reader) repo[T] {
	repo := repo[T]{}
	err := json.NewDecoder(in).Decode(&repo)
	if err != nil {
		panic(err)
	}

	return repo
}

var ErrNotFound = errors.New("not found")

func (r *repo[T]) Create(ctx context.Context, v T) (int64, error) {
	r.ID += 1
	setIDIfExists(&v, r.ID)
	r.Items[r.ID] = v

	return r.ID, nil
}

func (r *repo[T]) Read(ctx context.Context, ids []int64) ([]T, error) {
	items := make([]T, len(ids))

	for i, id := range ids {
		v, ok := r.Items[id]
		if !ok {
			return nil, ErrNotFound
		}
		items[i] = v
	}

	return items, nil
}

func (r *repo[T]) Update(ctx context.Context, id int64, v T) error {
	setIDIfExists(&v, id)
	r.Items[id] = v
	return nil
}

func (r *repo[T]) Delete(ctx context.Context, id int64) error {
	delete(r.Items, id)
	return nil
}

func (r *repo[T]) List(ctx context.Context) ([]T, error) {
	items := make([]T, 0, len(r.Items))
	for _, v := range r.Items {
		items = append(items, v)
	}

	return items, nil
}
