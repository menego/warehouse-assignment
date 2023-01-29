package middlewares

import (
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"warehouse-assignment/internal/pkg/common/structs"
	"warehouse-assignment/internal/pkg/database/products"
)

func FetchProducts(db *structs.DBStructure) iris.Handler {
	return func(ctx iris.Context) {

		catalog, err := products.ListProductCatalog(db)
		if err != nil {
			logrus.Errorf("FetchProducts: %v", err)
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		err = ctx.JSON(catalog)
		if err != nil {
			logrus.Errorf("FetchProducts: Failed to write response to the request: %v", err)
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		ctx.StatusCode(iris.StatusOK)
	}
}

func SellProducts(db *structs.DBStructure) iris.Handler {
	return func(ctx iris.Context) {

		productName := ctx.Params().Get("name")
		err := products.SellProduct(productName, db)
		if err != nil {
			logrus.Errorf("SellProducts: %v", err)
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		ctx.StatusCode(iris.StatusOK)
	}
}