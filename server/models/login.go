package models

import (
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
)

type Login struct {
	Email    string
	Password string
}

func (m *Login) Validate(o orm.Ormer) error {
	valid := validation.Validation{}
	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")

	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")

	valid.Email(m.Email, "email")
	u := User{Email: m.Email}
	u, err := u.FindByEmail(o)
	if err != nil {
		beego.Error(err)
	}
	if u.Id == 0 {
		err := valid.SetError("email", "User not found")
		if err != nil {
			beego.Error(err)
		}
	} else {
		passwordValid := password.Verify(m.Password, u.Salt, u.Password, nil)
		if passwordValid == false {
			err := valid.SetError("password", "Password is wrong")
			if err != nil {
				beego.Error(err)
			}
		}
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logger.LogModelFieldNotValidError(constants.ProductModel, err.Key, err.Message)
		}
		e := errors.New(fmt.Sprintf("Model %v is not valid", constants.ProductModel))
		return e
	}
	return nil
}
