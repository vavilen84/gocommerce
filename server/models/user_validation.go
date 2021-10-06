package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/validation"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
)

func (m *User) validateEmailAlreadyInUse(o orm.Ormer, valid *validation.Validation) {
	u := User{Email: m.Email}
	err := u.FindByEmail(o)
	if err != nil {
		if err != orm.ErrNoRows {
			logger.LogOrmerError(constants.UserModel, err)
		}
	} else {
		if (u.Id != 0) && (u.Id != m.Id) {
			err := valid.SetError("email", "Email is already in use")
			if err != nil {
				logger.LogError(err)
			}
		}
	}
}

func (m *User) ValidateUserExists(o orm.Ormer, valid *validation.Validation) {
	err := m.FindById(o)
	if err != nil {
		if err != orm.ErrNoRows {
			logger.LogOrmerError(constants.UserModel, err)
		} else {
			errMsg := fmt.Sprintf("User with id #%d does not exist", m.Id)
			err := valid.SetError("id", errMsg)
			if err != nil {
				logger.LogError(err)
			}
		}
	}
}

func (m *User) getValidationErrors(valid *validation.Validation) error {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logger.LogModelFieldNotValidError(constants.UserModel, err.Key, err.Message)
		}
		e := errors.New(fmt.Sprintf("Model %v is not valid", constants.UserModel))
		return e
	}
	return nil
}

func (m *User) validateCommonFields(valid *validation.Validation) {
	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")
	valid.Email(m.Email, "email")

	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")

	valid.Required(m.Salt, "salt")

	valid.Required(m.Role, "salt")
	valid.Range(m.Role, UserRoleCustomer, UserRoleAdmin, "role")

	valid.Required(m.FirstName, "first_name")
	valid.MaxSize(m.FirstName, 255, "first_name")

	valid.Required(m.LastName, "last_name")
	valid.MaxSize(m.LastName, 255, "last_name")
}

func (m *User) validateOnUpdate(o orm.Ormer) error {
	valid := validation.Validation{}
	valid.Required(m.Id, "id")
	m.ValidateUserExists(o, &valid)
	m.validateCommonFields(&valid)
	m.validateEmailAlreadyInUse(o, &valid)
	err := m.getValidationErrors(&valid)
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (m *User) validateOnInsert(o orm.Ormer) error {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.validateEmailAlreadyInUse(o, &valid)
	err := m.getValidationErrors(&valid)
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
