package products

import (
	"errors"
	"fmt"
	"math"
	"warehouse-assignment/internal/pkg/common/structs"
	"warehouse-assignment/internal/pkg/database/inventory"
)

var EmptyProduct = structs.Product{
	Name:       "",
	Price:      0,
	Components: []structs.ProductComponent{},
}

func isEmpty(p structs.Product) bool {
	return p.Name == "" && p.Price == 0 && (p.Components == nil || len(p.Components) == 0)
}

func ListProductCatalog(memDb *structs.DBStructure) ([]structs.ProductAvailability, error) {
	catalog := []structs.ProductAvailability{}

	for pName, quantity := range memDb.ProductQuantities {

		product, err := GetProduct(pName, memDb)
		if err != nil {
			return []structs.ProductAvailability{}, fmt.Errorf("ListProductCatalog: %w", err)
		}

		catalog = append(catalog, structs.ProductAvailability{
			Product:  product,
			Availability: quantity,
		})
	}

	return catalog, nil
}

func SaveProduct(p structs.Product, memDb *structs.DBStructure) {
	memDb.ProductBlueprints[p.Name] = p
}

func UpdateProductAvailability(name string, memDb *structs.DBStructure) {
	memDb.ProductQuantities[name] = computeAvailability(name, memDb)
}

func updateProductsAvailabilities(memDb *structs.DBStructure) {
	for _, prodBlueprint := range memDb.ProductBlueprints {
		UpdateProductAvailability(prodBlueprint.Name, memDb)
	}
}

func GetProduct(name string, memDb *structs.DBStructure) (structs.Product, error) {

	if isEmpty(memDb.ProductBlueprints[name]) {
		return EmptyProduct, errors.New("GetProduct: product does not exist")
	}
	return memDb.ProductBlueprints[name], nil
}

func SellProduct(name string, memDb *structs.DBStructure) error {
	if isEmpty(memDb.ProductBlueprints[name]) {
		return errors.New("GetProduct: product does not exist")
	}
	if memDb.ProductQuantities[name] <= 0 {
		return errors.New("SellProduct: not enough stock components")
	}

	for _, component := range memDb.ProductBlueprints[name].Components {
		invArticle := inventory.GetArticle(component.ArticleId, memDb)

		invArticle.Stock -= component.Quantity
		inventory.SaveArticle(invArticle, memDb)
	}

	updateProductsAvailabilities(memDb)

	return nil
}

func computeAvailability(name string, memDb *structs.DBStructure) int {
	productBlueprint := memDb.ProductBlueprints[name]
	currentAvailability := math.MaxInt32
	for _, component := range productBlueprint.Components {
		quantityBasedOnComponent := memDb.Inventory[component.ArticleId].Stock / component.Quantity
		if quantityBasedOnComponent < currentAvailability {
			currentAvailability = quantityBasedOnComponent
		}
	}

	return currentAvailability
}
