package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/validation"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
	"regexp"
)

type Product struct {
	Id    int64  `json:"id" column:"id"`
	Title string `json:"title" column:"title"`
	SKU   string `json:"sku" column:"sku"`
	Price int    `json:"price" column:"price"`

	CreatedAt int `json:"created_at" column:"created_at"`
	UpdatedAt int `json:"updated_at" column:"updated_at"`
	DeletedAt int `json:"deleted_at" column:"deleted_at"`
}

func (p Product) validateOnInsert() bool {
	valid := validation.Validation{}
	valid.Required(p.Title, "title_required")
	valid.Required(p.SKU, "sku_required")
	valid.Match(p.SKU, regexp.MustCompile(`^[a-z0-9_-]*$`), "sku_match")
	valid.Required(p.Price, "price_required")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logger.LogModelFieldNotValidError(constants.ProductModel, err.Key, err.Message)
		}
		return false
	}
	return true
}

func (p *Product) Insert(o orm.Ormer) error {
	valid := p.validateOnInsert()
	if !valid {
		logger.LogModelNotValidError(constants.ProductModel)
	}
	_, err := o.Insert(p)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	return nil
}
