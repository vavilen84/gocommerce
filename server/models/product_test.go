package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/store"
	"github.com/vavilen84/gocommerce/validation"
	"testing"
)

func TestProduct_ValidateOnCreate(t *testing.T) {
	m := Product{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, &m)
	assert.NotNil(t, err)
	assert.NotEmpty(t, err[constants.ProductPriceField])
	assert.NotEmpty(t, err[constants.ProductSKUField])
	assert.NotEmpty(t, err[constants.ProductTitleField])

	m = Product{
		Title: "product",
		SKU:   "sku-123_123",
		Price: 100,
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, &m)
	assert.NotNil(t, err)
}

func TestProduct_Create(t *testing.T) {
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	c := Product{
		Title: "title",
		SKU:   "sku",
		Price: 1,
	}
	err := c.Create(ctx, conn)
	assert.Nil(t, err)

	//c = Product{}
	//err = c.FindById(ctx, conn, id)
	//assert.Nil(t, err)
	//assert.Equal(t, c.Id, id)
	//assert.Equal(t, c.Name, name)
	//assert.Equal(t, c.Capacity, &capacity)
}
