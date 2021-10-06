package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
)

func (m *User) Insert(o orm.Ormer) (err error) {
	m.setPassword()
	m.setTimestampsOnCreate()
	err = m.validateOnInsert(o)
	if err != nil {
		logger.LogError(err)
		return err
	}
	_, err = o.Insert(m)
	if err != nil {
		logger.LogOrmerError(constants.UserModel, err)
		return err
	}
	return nil
}

func (m *User) Update(o orm.Ormer) (err error) {
	m.setPasswordOnUpdate(o)
	m.setTimestampsOnUpdate()
	err = m.validateOnUpdate(o)
	if err != nil {
		logger.LogError(err)
		return err
	}
	_, err = o.Update(m)
	if err != nil {
		logger.LogOrmerError(constants.UserModel, err)
	}
	return
}
