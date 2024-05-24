package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zoomea/meal-planning-app/biz"
)

func getManyRecipes(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	recipes, err := biz.ListRecipes(ctx, state.recipeDB)
	if err != nil {
		return nil, 500, err
	}
	return recipes, 200, nil
}

func getRecipe(req *http.Request, state State) (any, int, error) {
	idStr := req.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, 400, fmt.Errorf("Could not parse '%s' as integer", idStr)
	}

	ctx := req.Context()

	recipes, err := biz.GetRecipes(ctx, state.recipeDB, []int64{id})
	if err != nil {
		return nil, 500, err
	}
	if len(recipes) == 0 {
		return nil, 404, fmt.Errorf("no recipe with id %d", id)
	}

	return recipes[0], 200, nil
}

func addRecipe(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	var rec biz.Recipe
	if err := json.NewDecoder(req.Body).Decode(&rec); err != nil {
		return nil, 400, err
	}

	id, err := biz.AddRecipe(ctx, state.recipeDB, rec)
	if err != nil {
		return nil, 500, err
	}
	return id, 200, nil
}

func deleteRecipe(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	idStr := req.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, 400, fmt.Errorf("Could not parse '%s' as integer", idStr)
	}

	err = biz.DeleteRecipe(ctx, state.recipeDB, id)
	if err != nil {
		return nil, 500, err
	}
	return nil, 200, nil
}
