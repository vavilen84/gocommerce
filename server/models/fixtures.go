package models

const (
	Product1Sku = "product_1_sku"
)

var (
	productsFixtures map[int]Product
)

func init() {
	productsFixtures = map[int]Product{
		1: {
			Title: "Product #1 title",
			SKU:   Product1Sku,
			Price: 1,
		},
	}
}
