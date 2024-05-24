package biz

import (
	"context"
	"slices"

	"github.com/Zoomea/meal-planning-app/db"
)

// Returns the total list of ingredients required to cook all the dishes in the given time
// period as a map keyed by the ID of the ingredient.
func CalculateIngredients(ctx context.Context, sDB db.ScheduleStore, rDB db.Crudler[Recipe], start, end db.Date) (map[int64]IngredQuant, error) {
	scheds, err := sDB.List(ctx, start, end)
	if err != nil {
		return nil, err
	}

	// TODO room to preallocate array
	var recipeIDs []int64
	for _, sched := range scheds {
		recipeIDs = append(recipeIDs, sched.Recipes...)
	}

	// Deduplicate
	slices.Sort(recipeIDs)
	recipeIDs = slices.Compact(recipeIDs)

	recipes, err := rDB.Read(ctx, recipeIDs)
	if err != nil {
		return nil, err
	}

	// Aggregate the ingredient lists of all the recipes together
	shoppingList := make(map[int64]IngredQuant)
	for _, recipe := range recipes {
		for _, ingr := range recipe.Ingredients {
			s, ok := shoppingList[ingr.Ingredient.ID]
			if !ok {
				s = ingr
			} else {
				s.Quantity += ingr.Quantity
			}

			shoppingList[ingr.Ingredient.ID] = s
		}
	}

	return shoppingList, nil
}
