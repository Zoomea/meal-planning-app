package biz

type IngredQuant struct {
	// Embedded to flatten JSON
	Ingredient
	Quantity int64
	Unit     string
}

type Ingredient struct {
	ID   int64
	Name string
}
