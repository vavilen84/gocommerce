package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
	model := Product{
		Title: "title",
		SKU:   "sku",
		Price: 1,
	}
	err := model.Insert(o)
	assert.Nil(t, err)
}
