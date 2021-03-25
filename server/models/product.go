package models

type Product struct {
	Id    uint32 `json:"id" column:"id" validate:"required"`
	Title string `json:"title" column:"title" validate:"required"`
	SKU   string `json:"sku" column:"sku" validate:"required"`
}
