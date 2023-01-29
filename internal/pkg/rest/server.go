package rest

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/sirupsen/logrus"
	"time"
	"warehouse-assignment/internal/pkg/common/structs"
	"warehouse-assignment/internal/pkg/rest/middlewares"
)

var (
	appLog          *logrus.Entry
)

func Start(db *structs.DBStructure){
	app := iris.New()
	app.UseRouter(cors.New().AllowOrigin("*").Handler())
	v1 := app.Party("/api/v1")

	v1.Get("/products",
		middlewares.FetchProducts(db),
	)

	v1.Put("/products/{name}/sell",
		middlewares.SellProducts(db),
	)

	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(manageGracefulShutdown(app, idleConnsClosed))

	// iris.WithoutPathCorrection option removes the trailing slash matching, so /api/whatever will NOT be the same as /api/whatever/ .See: https://docs.iris-go.com/iris/contents/routing#behavior
	err := app.Listen(":3000", iris.WithoutBodyConsumptionOnUnmarshal, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed), iris.WithoutPathCorrection)
	<-idleConnsClosed
	if err != nil {
		appLog.Panicf("Error while starting http server, %v", err)
	}
}

//manageGracefulShutdown this method allows closing the various resources gracefully (e.g. db connection, cache etc.)
// before shutting down the server
func manageGracefulShutdown(app *iris.Application, idleConnsClosed chan struct{}) func() {
	return func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := app.Shutdown(ctx); err != nil {
			appLog.Errorf("Error while shutting down http server, %v", err)
		}
		close(idleConnsClosed)
	}
}