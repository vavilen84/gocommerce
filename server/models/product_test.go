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

	setTestAppEnv()
	store.InitTestDB()
	conn, ctx := store.GetNewTestDBConn()
	defer conn.Close()
	prepareTestDB(ctx, conn)

	model := Product{
		Title: "title",
		SKU:   "sku",
		Price: 1,
	}
	err := model.Create(ctx, conn)
	assert.Nil(t, err)

	modelFromDb, err := FindProductById(ctx, conn, model.GetId())
	assert.Nil(t, err)
	assert.Equal(t, model.Id, modelFromDb.Id)
	assert.Equal(t, model.Title, modelFromDb.Title)
	assert.Equal(t, model.SKU, modelFromDb.SKU)
	assert.Equal(t, model.Price, modelFromDb.Price)
}
