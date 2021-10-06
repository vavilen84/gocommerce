package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
)

const (
	_ = iota
	UserRoleCustomer
	UserRoleAdmin
)

type User struct {
	Id        int    `json:"id" orm:"auto"`
	Email     string `json:"email" orm:"column(email);unique"`
	Password  string `json:"password" orm:"column(password)"`
	Salt      string `json:"salt" orm:"column(salt)"`
	Role      int    `json:"role" orm:"column(role)"`
	FirstName string `json:"first_name" orm:"column(first_name)"`
	LastName  string `json:"last_name" orm:"column(last_name)"`

	CreatedAt int  `json:"created_at" orm:"column(created_at)"`
	UpdatedAt int  `json:"updated_at" orm:"column(updated_at)"`
	DeletedAt *int `json:"deleted_at" orm:"column(deleted_at)"`
}

func (m *User) FindByEmail(o orm.Ormer) (err error) {
	err = o.QueryTable(constants.UserModel).Filter("email", m.Email).One(m)
	if err != nil {
		logger.LogOrmerError(constants.UserModel, err)
	}
	return
}

func (m *User) FindById(o orm.Ormer) (err error) {
	err = o.QueryTable(constants.UserModel).Filter("id", m.Id).One(m)
	if err != nil {
		logger.LogOrmerError(constants.UserModel, err)
	}
	return
}
