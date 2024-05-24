package web_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/testutil"
	"github.com/Zoomea/meal-planning-app/web"
)

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

var (
	port    = 8080
	baseURL = fmt.Sprintf("http://127.0.0.1:%d", port)
)

func TestRecipe(t *testing.T) {
	ready := make(chan struct{})
	go func() {
		err := web.Serve("./../public/", 8080, ready)
		fmt.Printf("err: %+v", err)
	}()

	// Wait until the server is listening to the socket
	<-ready

	recipes := getRecipes(t)
	testutil.Assert(t, len(recipes), 0)

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

	recipe := getRecipe(t, id)
	testutil.Assert(t, recipe, newRecipe)
}

func getRecipes(t *testing.T) []biz.Recipe {
	r, err := http.Get(baseURL + "/api/recipe/")
	testutil.Assert(t, err, nil)

	var res Response[[]biz.Recipe]
	err = json.NewDecoder(r.Body).Decode(&res)
	testutil.Assert(t, err, nil)

	return res.Data
}

func getRecipe(t *testing.T, id int64) biz.Recipe {
	url := fmt.Sprintf(baseURL+"/api/recipe/%d", id)
	r, err := http.Get(url)
	testutil.Assert(t, err, nil)

	var res Response[biz.Recipe]
	err = json.NewDecoder(r.Body).Decode(&res)
	testutil.Assert(t, err, nil)

	return res.Data
}

func addRecipe(t *testing.T, recipe biz.Recipe) int64 {
	body, _ := json.Marshal(recipe)
	r, err := http.Post(baseURL+"/api/recipe/", "application/json", bytes.NewReader(body))

	testutil.Assert(t, r.StatusCode, 200)
	testutil.Assert(t, err, nil)
	var resp Response[int64]

	err = json.NewDecoder(r.Body).Decode(&resp)
	testutil.Assert(t, err, nil)

	return resp.Data
}
