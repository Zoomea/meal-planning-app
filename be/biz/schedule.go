package biz

import (
	"context"
	"slices"

	"github.com/Zoomea/meal-planning-app/db"
)

// Schedule contains pointers to recipes
type Schedule struct {
	Date    db.Date           `json:"date"`
	Type    string            `json:"type"`
	Recipes []BasicRecipeInfo `json:"recipes"`
}

type BasicRecipeInfo struct {
	Name string
}

func ListSchedules(ctx context.Context, rDB db.Crudler[Recipe], db db.ScheduleStore, start, end db.Date) ([]Schedule, error) {
	scheds, err := db.List(ctx, start, end)
	if err != nil {
		return nil, err
	}

	enrichedScheds := make([]Schedule, len(scheds))
	for i, sched := range scheds {
		enrichedScheds[i], err = enrichSched(ctx, rDB, sched)
		if err != nil {
			return nil, err
		}
	}

	return enrichedScheds, nil
}

// Given a schedule, it goes and fetchs the recipes the schedule contains
// and inserts them into the schedule.
func enrichSched(ctx context.Context, db db.Crudler[Recipe], sch db.Schedule) (Schedule, error) {
	recipes, err := db.Read(ctx, sch.Recipes)
	if err != nil {
		return Schedule{}, err
	}

	enrichedSched := Schedule{
		Date:    sch.Date,
		Type:    sch.Type,
		Recipes: make([]BasicRecipeInfo, len(recipes)),
	}

	for i, recipe := range recipes {
		enrichedSched.Recipes[i] = BasicRecipeInfo{Name: recipe.Name}
	}

	return enrichedSched, nil
}

func AddSchedule(ctx context.Context, conn db.ScheduleStore, sch db.Schedule) error {
	return conn.UpsertSchedule(ctx, sch)

}

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
