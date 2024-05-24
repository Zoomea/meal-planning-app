package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Zoomea/meal-planning-app/biz"
)

func getShoppingList(req *http.Request, state State) (any, int, error) {
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

	recipes, err := biz.CalculateIngredients(ctx, state.scheduleDB, state.recipeDB, from, to)
	if err != nil {
		return nil, 500, err
	}

	return recipes, 200, nil
}
