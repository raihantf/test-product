package result

type AddProductResult struct {
	Id int `json:"id"`
}

type GetLastInsertCategoryIdResult struct {
	Category_Id int `json:"category_id"`
}

type GetProductCategories struct {
	ProductCategoriesName string `json:"product_categories_name"`
}

type Products struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Category_Id int    `json:"category_id"`
}
