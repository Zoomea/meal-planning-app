package web_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/db"
	"github.com/Zoomea/meal-planning-app/testutil"
	"github.com/Zoomea/meal-planning-app/web"
)

func TestShoppingList(t *testing.T) {
	ready := make(chan struct{})
	go func() {
		err := web.Serve("./../public/", 8080, ready)
		fmt.Printf("err: %+v", err)
	}()

	// Wait until the server is listening to the socket
	<-ready

	newRecipe := biz.Recipe{
		Name:         "Test Recipe1",
		Description:  "Test Description",
		PrepTimeMins: 80,
		CookTimeMins: 100,
		Difficulty:   5,
		Ingredients: []biz.IngredQuant{{
			Ingredient: biz.Ingredient{
				ID:   1,
				Name: "cumin",
			},
			Quantity: 10,
			Unit:     "gram",
		},
			{
				Ingredient: biz.Ingredient{
					ID:   2,
					Name: "tomato",
				},
				Quantity: 2,
				Unit:     "",
			},
		},
	}

	id := addRecipe(t, newRecipe)
	newRecipe.ID = id

	date := db.Date{
		Day:   1,
		Month: 2,
		Year:  2023,
	}

	schedule := db.Schedule{
		Date:    date,
		Type:    "breakfast",
		Recipes: []int64{id},
	}

	addSchedule(t, schedule)

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

	list := getShoppingList(t, from, to)
	testutil.Assert(t, len(list), 2)
}

func getShoppingList(t *testing.T, from, to db.Date) map[int64]biz.IngredQuant {
	queryParams := "from=" + formatDate(from) + "&to=" + formatDate(to)
	r, err := http.Get(baseURL + "/api/shopping-list/?" + queryParams)

	testutil.Assert(t, r.StatusCode, 200)
	testutil.Assert(t, err, nil)
	var resp Response[map[int64]biz.IngredQuant]

	err = json.NewDecoder(r.Body).Decode(&resp)
	testutil.Assert(t, err, nil)

	return resp.Data
}
