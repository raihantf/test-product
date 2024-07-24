package routes

import (
	product_controller "test-product/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Post("/product", product_controller.AddProduct)
	app.Post("/product-list", product_controller.GetListProduct)
}
