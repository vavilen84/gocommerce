package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
	"time"
)

func (m *User) setPasswordOnUpdate(o orm.Ormer) {
	oldUser, err := FindUserById(o, m.Id)
	if err != nil {
		logger.LogOrmerError(constants.UserModel, err)
	}
	if m.Password != "" {
		m.setPassword()
	} else {
		m.Password = oldUser.Password
		m.Salt = oldUser.Salt
	}
}

func (m *User) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func (m *User) setTimestampsOnUpdate() {
	now := int(time.Now().Unix())
	m.UpdatedAt = now
}

func (m *User) setPassword() {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.Salt = salt
}
