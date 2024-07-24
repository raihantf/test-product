package controllers

import (
	product_request "test-product/app/dto/request"
	product_response "test-product/app/dto/response"
	product_service "test-product/app/services"

	"github.com/gofiber/fiber/v2"
)

// @Tags Product
// @Summary Add Product
// @Router /product [post]
// @Param request body product_request.AddProductRequest true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Success 200 {string} string "message": "Sucessfully Add New Product"
// @Failure 400
func AddProduct(ctx *fiber.Ctx) error {
	data := new(product_request.AddProductRequest)

	if err := ctx.BodyParser(data); err != nil {
		return err
	}

	result, err := product_service.AddProductService(data)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(&fiber.Map{
		"message": result,
	})
}

// @Tags Product
// @Summary Get List Product
// @Router /product-list [post]
// @Param request body product_request.GetListRequest true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Success 200 {array} product_response.Products
// @Failure 400
func GetListProduct(ctx *fiber.Ctx) error {
	data := new(product_request.GetListRequest)

	if err := ctx.BodyParser(data); err != nil {
		return err
	}

	var result []*product_response.Products

	result, err := product_service.GetListProductService(data)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(&fiber.Map{
		"message": result,
	})
}
