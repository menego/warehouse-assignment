package main

import (
	"log"
	"warehouse-assignment/internal/pkg/database"
	"warehouse-assignment/internal/pkg/rest"
)

func init() {

}

func main(){
	err := database.InitData()
	if err != nil {
		log.Panic(err)
	}

	rest.Start(database.GetDbConnection())
	//catalog, err := products.ListProductCatalog(database.GetDbConnection())
}