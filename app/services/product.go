package services

import (
	database "test-product/app/database"
	product_request "test-product/app/dto/request"
	product_response "test-product/app/dto/response"
	common_error "test-product/app/error"
	product_repository "test-product/app/repository"
)

func AddProductService(data *product_request.AddProductRequest) (resp string, err error) {
	db, err := database.DBInit()
	if err != nil {
		panic(err)
	}

	getProductCategory := product_repository.GetProductCategories(db.Conn, data.Category_Id)
	if getProductCategory == "" {
		panic(&common_error.BadInputError{
			Message: "Product category doesn't exist",
		})
	}

	addProduct := product_repository.AddProductRepository(db.Conn, data.ProductName, data.Description, data.Category_Id)
	if addProduct < 1 {
		panic(&common_error.ApiError{
			Message: "Failed insert new product",
		})
	}

	resp = "Sucessfully Add New Product"

	return resp, err
}

func GetListProductService(data *product_request.GetListRequest) (resp []*product_response.Products, err error) {
	db, err := database.DBInit()
	if err != nil {
		panic(err)
	}

	getListProduct := product_repository.GetListProductRepository(db.Conn, data)

	for i := 0; i < len(getListProduct); i++ {
		resp = append(resp, getListProduct[i])
	}

	return resp, err
}
