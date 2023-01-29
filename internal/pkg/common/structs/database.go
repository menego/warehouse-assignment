package structs

type DBStructure struct {
	Inventory         map[int]Article
	ProductBlueprints map[string]Product
	ProductQuantities map[string]int
}
