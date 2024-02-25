package biz

type IngredQuant struct {
	// Embedded to flatten JSON
	Ingredient
	Quantity int64  `json:"quantity"`
	Unit     string `json:"unit"`
}

type Ingredient struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
