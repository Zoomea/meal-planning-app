package db_test

import (
	"testing"

	"github.com/Zoomea/meal-planning-app/db"
)

func TestDB(t *testing.T) {
	d := db.NewDB()
	r := d.Repo("test")

	testStr := "this is a test"
	id := r.Add([]byte(testStr))

	data, err := r.Get(id)
	if err != nil {
		t.Logf(err.Error())
	}

	assert(t, testStr, string(data))
}

func assert[T comparable](t *testing.T, got, expected T) {
	t.Helper()
	if expected != got {
		t.Fatalf(`
-- expected --
%v
-- got --
%v`, expected, got)
	}
}
