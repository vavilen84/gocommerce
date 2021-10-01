package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

func init() {
	orm.RegisterModel(new(Product))
}
