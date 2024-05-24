package web_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/db"
	"github.com/Zoomea/meal-planning-app/testutil"
	"github.com/Zoomea/meal-planning-app/web"
)

func TestSchedule(t *testing.T) {
	ready := make(chan struct{})
	go func() {
		err := web.Serve("./../public/", 8080, ready)
		fmt.Printf("err: %+v", err)
	}()

	// Wait until the server is listening to the socket
	<-ready

	from := db.Date{
		Day:   1,
		Month: 2,
		Year:  2023,
	}
	to := db.Date{
		Day:   2,
		Month: 2,
		Year:  2023,
	}

	schedules := listSchedules(t, from, to)
	testutil.Assert(t, len(schedules), 0)

	date := db.Date{
		Day:   1,
		Month: 2,
		Year:  2023,
	}

	schedule := db.Schedule{
		Date:    date,
		Type:    "breakfast",
		Recipes: []int64{1},
	}

	addSchedule(t, schedule)

	schedules = listSchedules(t, from, to)
	testutil.Assert(t, len(schedules), 1)
}

func listSchedules(t *testing.T, from, to db.Date) []biz.Schedule {
	queryParams := "from=" + formatDate(from) + "&to=" + formatDate(to)
	r, err := http.Get(baseURL + "/api/schedule/?" + queryParams)
	testutil.Assert(t, err, nil)

	var res Response[[]biz.Schedule]
	err = json.NewDecoder(r.Body).Decode(&res)
	testutil.Assert(t, err, nil)

	return res.Data
}

func formatDate(date db.Date) string {
	return fmt.Sprintf("%d-%d-%d", date.Year, date.Month, date.Day)
}

// TODO ideally I shouldn't be using the definition from the DB
func addSchedule(t *testing.T, schedule db.Schedule) {
	body, _ := json.Marshal(schedule)
	r, err := http.Post(baseURL+"/api/schedule/", "application/json", bytes.NewReader(body))

	testutil.Assert(t, r.StatusCode, 200)
	testutil.Assert(t, err, nil)
	var resp Response[json.RawMessage]

	err = json.NewDecoder(r.Body).Decode(&resp)
	testutil.Assert(t, err, nil)
}
