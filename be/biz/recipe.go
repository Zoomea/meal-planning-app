package biz

import "context"

type Conf struct {
	dbConn RecipeRepo
}

type Recipe struct {
	ID   int64  `json:"id"`
	Name string `json:"recipe"`
}

// Note that we expect the same Recipe struct from the DB that we will later
// pass to the user which is a leaky abstraction.
// Change it later if it causes problems
type RecipeRepo interface {
	Get(context.Context, int64) (Recipe, error)
	GetAll(context.Context) ([]Recipe, error)
	Add(context.Context, Recipe) (int64, error)
	Delete(context.Context, int64) error
}

func GetRecipe(ctx context.Context, conf Conf, id int64) (Recipe, error) {
	return conf.dbConn.Get(ctx, id)
}

func GetRecipes(ctx context.Context, conf Conf) ([]Recipe, error) {
	return conf.dbConn.GetAll(ctx)
}

func AddRecipe(ctx context.Context, conf Conf, recipe Recipe) (int64, error) {
	return conf.dbConn.Add(ctx, recipe)
}

func DeleteRecipe(ctx context.Context, conf Conf, id int64) error {
	return conf.dbConn.Delete(ctx, id)
}
