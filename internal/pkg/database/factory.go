package database

import (
	"fmt"
	"warehouse-assignment/internal/pkg/common/structs"
	"warehouse-assignment/internal/pkg/database/inventory"
	"warehouse-assignment/internal/pkg/database/products"
	file "warehouse-assignment/internal/pkg/readers/json"
)


var memDb = structs.DBStructure{}

// InitData is a convenience method to load the provided files in the asset folder into the memory database
func InitData() error {

	memDb.Inventory = make(map[int]structs.Article)
	memDb.ProductBlueprints = make(map[string]structs.Product)
	memDb.ProductQuantities = make(map[string]int)

	fileReader := file.Reader{}
	articles, err := file.ReadInventory("./assets/inventory.json", &fileReader)
	if err != nil {
		return fmt.Errorf("InitData: %w", err)
	}
	for _, article := range articles {

		inventory.SaveArticle(article, &memDb)
	}


	pBlueprints, err := file.ReadProducts("./assets/products.json", &fileReader)
	if err != nil {
		return fmt.Errorf("InitData: %w", err)
	}

	for _, prodBlueprint := range pBlueprints {
		products.SaveProduct(prodBlueprint, &memDb)
		products.UpdateProductAvailability(prodBlueprint.Name, &memDb)
	}

	return nil
}

// GetDbConnection method that mimics a hypothetical db interface that is passed to access the data layer
func GetDbConnection() *structs.DBStructure {
	return &memDb
}
