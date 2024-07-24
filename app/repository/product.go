package repository

import (
	product_request "test-product/app/dto/request"
	product_response "test-product/app/dto/response"
	product_result "test-product/app/dto/result"

	"gorm.io/gorm"
)

func GetProductCategories(tx *gorm.DB, category_id int) string {
	var result string
	err := tx.Raw(`SELECT product_categories_name FROM product_categories WHERE id = ?`, category_id).Scan(&result).Error
	if err != nil {
		panic(err)
	}

	return result
}

func AddProductRepository(tx *gorm.DB, product_name, description string, category_id int) int64 {
	product := &product_result.Products{
		ProductName: product_name,
		Description: description,
		Category_Id: category_id,
	}
	result := tx.Select("product_name", "description", "category_id").Create(&product)

	return result.RowsAffected
}

func GetListProductRepository(tx *gorm.DB, data *product_request.GetListRequest) []*product_response.Products {
	var products []*product_response.Products

	if data.Keyword != "" && data.Category_Id > 0 {
		tx.Raw("SELECT * FROM products WHERE product_name LIKE '%"+data.Keyword+"%' OR description LIKE '%"+data.Keyword+"%' AND category_id = ?", data.Category_Id).Scan(&products)
	} else if data.Keyword != "" {
		tx.Raw("SELECT * FROM products WHERE product_name LIKE '%" + data.Keyword + "%' OR description LIKE '%" + data.Keyword + "%'").Scan(&products)
	} else if data.Category_Id > 0 {
		tx.Raw(`SELECT * FROM products WHERE category_id = ?`, data.Category_Id).Scan(&products)
	} else {
		tx.Raw(`SELECT * FROM products`).Scan(&products)
	}

	return products
}
