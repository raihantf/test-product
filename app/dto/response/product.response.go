package response

type AddProductResponse struct {
	ProductName           string `json:"product_name"`
	Description           string `json:"description"`
	ProductCategoriesName string `json:"product_categories"`
}

type Products struct {
	ID          int    `gorm:"primarykey"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}
