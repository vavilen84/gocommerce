package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createProduct(t *testing.T, o orm.Ormer) *Product {
	p := productsFixtures[1]
	err := p.Insert(o)
	assert.Nil(t, err)
	return &p
}

func TestProduct_Create(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
	createProduct(t, o)
}

func TestProduct_FindById(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
	p := createProduct(t, o)

	modelFromDb := Product{Id: p.Id}
	err := modelFromDb.FindById(o)
	assert.Nil(t, err)
	assert.Equal(t, p.Title, modelFromDb.Title)
	assert.Equal(t, p.SKU, modelFromDb.SKU)
	assert.Equal(t, p.Price, modelFromDb.Price)
}

func TestProduct_FindByDKU(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
	p := createProduct(t, o)

	modelFromDb := Product{SKU: p.SKU}
	err := modelFromDb.FindBySKU(o)
	assert.Nil(t, err)
	assert.Equal(t, p.Title, modelFromDb.Title)
	assert.Equal(t, p.SKU, modelFromDb.SKU)
	assert.Equal(t, p.Price, modelFromDb.Price)
}

func TestProduct_Update(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
	p := createProduct(t, o)
	newTitle := "New Title"
	p.Title = newTitle
	err := p.Update(o)
	assert.Nil(t, err)

	modelFromDb := Product{Id: p.Id}
	err = modelFromDb.FindById(o)
	assert.Nil(t, err)
	assert.Equal(t, newTitle, modelFromDb.Title)
}

func TestProduct_Delete(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
	p := createProduct(t, o)

	modelFromDb := Product{Id: p.Id}
	err := modelFromDb.FindById(o)
	assert.Nil(t, err)

	err = modelFromDb.Delete(o)
	assert.Nil(t, err)

	deletedProduct := Product{Id: p.Id}
	err = deletedProduct.FindById(o)
	assert.NotNil(t, err)
	assert.Equal(t, orm.ErrNoRows, err)
}
