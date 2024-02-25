package biz_test

import (
	"context"
	"testing"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/db/fsdatabase"
	"github.com/Zoomea/meal-planning-app/testutil"
)

func TestBasicRecipes(t *testing.T) {
	recipeDB := fsdatabase.New[biz.Recipe]()

	ctx := context.Background()

	cumin := biz.Ingredient{
		ID:   1,
		Name: "cumin",
	}

	recipe := biz.Recipe{
		Name: "test 1",
		Ingredients: []biz.IngredQuant{{
			Ingredient: cumin,
			Quantity:   120,
			Unit:       "grams",
		}},
	}

	id, err := biz.AddRecipe(ctx, recipeDB, recipe)
	testutil.Assert(t, err, nil)

	recipes, err := biz.GetRecipes(ctx, recipeDB, []int64{id})
	testutil.Assert(t, err, nil)

	t.Logf("recipes: %+v", recipes)

}
