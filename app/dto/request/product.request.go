package request

import validation "github.com/go-ozzo/ozzo-validation"

type AddProductRequest struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Category_Id int    `json:"category_id"`
}

func (dto *AddProductRequest) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ProductName, validation.Length(3, 50), validation.Required),
		validation.Field(&dto.Description, validation.Length(6, 12), validation.Required),
		validation.Field(&dto.Category_Id, validation.Min(1), validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetListRequest struct {
	Keyword     string `json:"keyword"`
	Category_Id int    `json:"category_id"`
}

func (dto *GetListRequest) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Keyword, validation.Length(1, 50)),
		validation.Field(&dto.Category_Id, validation.Min(1)),
	); err != nil {
		return err
	}
	return nil
}
