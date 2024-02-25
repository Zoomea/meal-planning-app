package fsdatabase_test

import (
	"context"
	"strings"
	"testing"

	"github.com/Zoomea/meal-planning-app/db/fsdatabase"
	"github.com/Zoomea/meal-planning-app/testutil"
)

var ctx = context.Background()

type item struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func TestParsing(t *testing.T) {
	in := strings.NewReader(`{"id":2, "items": {"1":{"name":"item 1"}, "2":{"name":"item 2"}}}`)
	conn := fsdatabase.FromJSON[item](in)

	_, err := conn.List(ctx)
	testutil.Assert(t, err, nil)
}

func TestCrud(t *testing.T) {
	data := []item{
		{Name: "item 1"},
		{Name: "item 2"},
	}

	conn := fsdatabase.New[item]()

	var err error
	data[0].ID, err = conn.Create(ctx, data[0])
	testutil.Assert(t, err, nil)

	got, err := conn.List(ctx)
	testutil.Assert(t, err, nil)
	testutil.Assert(t, len(got), 1)
	testutil.Assert(t, got[0].Name, data[0].Name)

	got, err = conn.Read(ctx, []int64{data[0].ID})
	testutil.Assert(t, err, nil)
	testutil.Assert(t, len(got), 1)
	testutil.Assert(t, got[0].Name, data[0].Name)

	err = conn.Update(ctx, data[0].ID, data[1])
	testutil.Assert(t, err, nil)

	got, err = conn.Read(ctx, []int64{data[0].ID})
	testutil.Assert(t, err, nil)
	testutil.Assert(t, len(got), 1)
	testutil.Assert(t, got[0].Name, data[1].Name)
}
