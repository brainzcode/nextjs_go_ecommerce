package main

import (
	"github.com/anthdm/weavebox"
	"github.com/brainzcode/nextjs_go_ecommerce/api"
)

func main() {
	app := weavebox.New()

	adminRoute := app.Box("/admin")

	productHandler := &api.ProductHandler{}

	adminRoute.Get("/product", productHandler.HandleGetProduct)

	adminRoute.Post("/product", productHandler.HandlePostProduct)

	app.Serve(8000)
}
