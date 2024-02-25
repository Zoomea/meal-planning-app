package biz

import (
	"context"

	"github.com/Zoomea/meal-planning-app/db"
)

type Recipe struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	PrepTimeMins int64         `json:"prepTimeMins"`
	CookTimeMins int64         `json:"cookTimeMins"`
	Difficulty   int64         `json:"difficulty"`
	Ingredients  []IngredQuant `json:"ingredients"`
}

func GetRecipes(ctx context.Context, conn db.Crudler[Recipe], ids []int64) ([]Recipe, error) {
	return conn.Read(ctx, ids)
}

func ListRecipes(ctx context.Context, conn db.Crudler[Recipe]) ([]Recipe, error) {
	return conn.List(ctx)
}

func AddRecipe(ctx context.Context, conn db.Crudler[Recipe], recipe Recipe) (int64, error) {
	return conn.Create(ctx, recipe)
}

func DeleteRecipe(ctx context.Context, conn db.Crudler[Recipe], id int64) error {
	return conn.Delete(ctx, id)
}
