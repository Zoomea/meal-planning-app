package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/db"
)

func getManySchedules(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	fromStr := req.URL.Query().Get("from")
	toStr := req.URL.Query().Get("to")

	if len(fromStr) == 0 || len(toStr) == 0 {
		return nil, 400, errors.New(`must specify the range of dates you wish to query in the query parameters "from" and "to"`)
	}

	from, err := parseDate(fromStr)
	if err != nil {
		return nil, 400, fmt.Errorf("parsing 'from' as date: %w", err)
	}

	to, err := parseDate(toStr)
	if err != nil {
		return nil, 400, fmt.Errorf("parsing 'to' as date: %w", err)
	}

	recipes, err := biz.ListSchedules(ctx, state.recipeDB, state.scheduleDB, from, to)
	if err != nil {
		return nil, 500, err
	}
	return recipes, 200, nil
}

func parseDate(s string) (db.Date, error) {
	slice := strings.Split(s, "-")
	if len(slice) != 3 {
		return db.Date{}, errors.New("expected date in format YYYY-MM-DD")
	}

	year, err1 := strconv.ParseInt(slice[0], 10, 64)
	month, err2 := strconv.ParseInt(slice[1], 10, 64)
	day, err3 := strconv.ParseInt(slice[2], 10, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		return db.Date{}, errors.New("expected date in format YYYY-MM-DD")
	}

	return db.Date{Year: year, Month: month, Day: day}, nil

}

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}


func addSchedule(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	var sched db.Schedule
	if err := json.NewDecoder(req.Body).Decode(&sched); err != nil {
		return nil, 400, err
	}

	err := biz.AddSchedule(ctx, state.scheduleDB, sched)
	if err != nil {
		return nil, 500, err
	}
	return nil, 200, nil
}
