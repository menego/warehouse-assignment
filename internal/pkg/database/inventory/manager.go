package inventory

import "warehouse-assignment/internal/pkg/common/structs"

func GetArticle(id int, memDb *structs.DBStructure) structs.Article {
	return memDb.Inventory[id]
}

func SaveArticle(a structs.Article, memDb *structs.DBStructure) {
	memDb.Inventory[a.Id] = a
}