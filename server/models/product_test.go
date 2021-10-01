package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_Create(t *testing.T) {

	beforeTestRun()
	//db := store.GetTestDB()
	//ctx := store.GetDefaultDBContext()
	//conn, connErr := db.Conn(ctx)
	//if connErr != nil {
	//	logger.LogFatal(connErr)
	//}
	//defer conn.Close()
	prepareTestDB()
	o := orm.NewOrm()
	model := Product{
		Title: "title",
		SKU:   "sku",
		Price: 1,
	}
	err := model.Insert(o)
	assert.Nil(t, err)

	//modelFromDb, err := FindProductById(ctx, conn, model.GetId())
	//assert.Nil(t, err)
	//assert.Equal(t, model.Id, modelFromDb.Id)
	//assert.Equal(t, model.Title, modelFromDb.Title)
	//assert.Equal(t, model.SKU, modelFromDb.SKU)
	//assert.Equal(t, model.Price, modelFromDb.Price)
}
