package fixtures

import (
	"github.com/vavilen84/gocommerce/models"
)

var (
	products map[int]models.Product
)

func init() {
	products = map[int]models.Product{
		1: {
			Title: "Product #1 title",
			SKU:   "Product #1 SKU",
			Price: 1,
		},
	}
}

func GetProductFixtures() map[int]models.Product {
	return products
}
