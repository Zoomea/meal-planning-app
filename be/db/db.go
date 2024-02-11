package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type DB struct {
	repos map[string]*repository
}

func NewDB() DB {
	return DB{
		repos: make(map[string]*repository),
	}
}

func NewDBFromJSON(in io.Reader) DB {
	db := DB{}
	_ = json.NewDecoder(in).Decode(&db)
	return db
}

func (db *DB) Repo(name string) Repo {
	repo, ok := db.repos[name]
	if !ok {
		repo = &repository{
			id: 0,
			m:  make(map[int64][]byte),
		}
		db.repos[name] = repo
	}

	return repo
}

type Repo interface {
	Get(id int64) ([]byte, error)
	GetAll() ([]byte, error)
	Add(r []byte) int64
	Update(id int64, r []byte) error
	Delete(id int64) error
}

type repository struct {
	// The largest id in the list
	id int64
	// The actual data
	m map[int64][]byte
}

var ErrNotFound = errors.New("not found")

func (r *repository) Get(id int64) ([]byte, error) {
	item, ok := r.m[id]
	if !ok {
		return nil, fmt.Errorf("no item with id %d: %w", id, ErrNotFound)
	}

	return item, nil
}

func (r *repository) GetAll() ([]byte, error) {
	data, _ := json.Marshal(r.m) // theres simply no way this can fail here
	return data, nil
}

func (r *repository) Add(data []byte) int64 {
	r.id += 1
	r.m[r.id] = data
	return r.id
}

func (r *repository) Update(id int64, data []byte) error {
	_, ok := r.m[id]
	if !ok {
		return fmt.Errorf("not item with id %d: %w", id, ErrNotFound)
	}

	r.m[id] = data
	return nil
}

// Deleting non-existant entries is a no-op
func (r *repository) Delete(id int64) error {
	delete(r.m, id)
	return nil
}
