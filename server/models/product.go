package models

type Product struct {
	Id    uint32 `json:"id" column:"id"`
	Title string `json:"title"`
	SKU   string `json:"sku"`
}
