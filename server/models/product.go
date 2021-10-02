package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/validation"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
	"regexp"
	"time"
)

type Product struct {
	Id    int64  `json:"id" orm:"auto"`
	Title string `json:"title" orm:"column(title)"`
	SKU   string `json:"sku" orm:"column(sku)"`
	Price int64  `json:"price" orm:"column(price)"`

	CreatedAt *int64 `json:"created_at" orm:"column(created_at)"`
	UpdatedAt *int64 `json:"updated_at" orm:"column(updated_at)"`
	DeletedAt *int64 `json:"deleted_at" orm:"column(deleted_at)"`
}

func (p Product) validateOnInsert() error {
	valid := validation.Validation{}
	valid.Required(p.Title, "title_required")
	valid.Required(p.SKU, "sku_required")
	valid.Match(p.SKU, regexp.MustCompile(`^[a-z0-9_-]*$`), "sku_match")
	valid.Required(p.Price, "price_required")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logger.LogModelFieldNotValidError(constants.ProductModel, err.Key, err.Message)
		}
		e := errors.New(fmt.Sprintf("Model %v is not valid", constants.ProductModel))
		return e
	}
	return nil
}

func (p *Product) setTimestampsOnCreate() {
	now := time.Now().Unix()
	p.CreatedAt = &now
	p.UpdatedAt = &now
}

func (p *Product) Insert(o orm.Ormer) error {
	err := p.validateOnInsert()
	if err != nil {
		logger.LogError(err)
		return err
	}
	p.setTimestampsOnCreate()
	_, err = o.Insert(p)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	return nil
}

func (p *Product) FindById(o orm.Ormer) error {
	err := o.Read(p)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	return nil
}

func (p *Product) FindBySKU(o orm.Ormer) error {
	qs := o.QueryTable(p)
	err := qs.Filter("sku", p.SKU).One(p)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	return nil
}

func (p *Product) validateProductExists(o orm.Ormer) error {
	m := Product{Id: p.Id}
	err := o.Read(&m)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	return nil
}

func (p *Product) Update(o orm.Ormer) error {
	err := p.validateProductExists(o)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	_, err = o.Update(p)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	return nil
}

func (p *Product) Delete(o orm.Ormer) error {
	err := p.validateProductExists(o)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
		return err
	}
	_, err = o.Delete(p)
	if err != nil {
		logger.LogOrmerError(constants.ProductModel, err)
	}
	return nil
}
