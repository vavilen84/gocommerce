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
	Id    int    `json:"id" orm:"auto"`
	Title string `json:"title" orm:"column(title)"`
	SKU   string `json:"sku" orm:"column(sku)"`
	Price int    `json:"price" orm:"column(price)"`

	CreatedAt int  `json:"created_at" orm:"column(created_at)"`
	UpdatedAt int  `json:"updated_at" orm:"column(updated_at)"`
	DeletedAt *int `json:"deleted_at" orm:"column(deleted_at)"`
}

func (p Product) validateOnInsert() error {
	valid := validation.Validation{}

	valid.MaxSize(p.Title, 255, "title")
	valid.Required(p.Title, "title")

	valid.MaxSize(p.SKU, 255, "title")
	valid.Required(p.SKU, "sku")

	valid.Match(p.SKU, regexp.MustCompile(`^[a-z0-9_-]*$`), "sku")

	valid.Required(p.Price, "price")

	valid.Required(p.CreatedAt, "created_at")
	valid.Required(p.UpdatedAt, "updated_at")

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
	now := int(time.Now().Unix())
	p.CreatedAt = now
	p.UpdatedAt = now
}

func (p *Product) Insert(o orm.Ormer) error {
	p.setTimestampsOnCreate()
	err := p.validateOnInsert()
	if err != nil {
		logger.LogError(err)
		return err
	}
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
