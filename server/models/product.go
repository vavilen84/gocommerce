package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/database"
	"github.com/vavilen84/gocommerce/helpers"
	"github.com/vavilen84/gocommerce/validation"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"regexp"
)

type Product struct {
	Id    uint32 `json:"id" column:"id"`
	Title string `json:"title" column:"title"`
	SKU   string `json:"sku" column:"sku"`
	Price uint64 `json:"price" column:"price"`

	CreatedAt int64 `json:"created_at" column:"created_at"`
	UpdatedAt int64 `json:"updated_at" column:"updated_at"`
	DeletedAt int64 `json:"deleted_at" column:"deleted_at"`
}

func (m Product) GetId() uint32 {
	return m.Id
}

func (Product) GetTableName() string {
	return constants.ProductDBTable
}

func (Product) getValidationRules() validation.ScenarioRules {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			constants.ProductTitleField: "required,min=1,max=255",
			constants.ProductSKUField:   "required,min=1,max=255,sku",
			constants.ProductPriceField: "required,min=0,max=999999999999",
		},
	}
}

func (Product) getValidator() (v *validator.Validate) {
	v = validator.New()
	err := v.RegisterValidation("sku", ValidateSKU)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func ValidateSKU(fl validator.FieldLevel) (r bool) {
	pattern := `^[a-z0-9_-]*$`
	r, err := regexp.Match(pattern, []byte(fl.Field().String()))
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (m Product) Create(ctx context.Context, conn *sql.Conn) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
	if err != nil {
		log.Println(err)
		return
	}
	err = database.Insert(ctx, conn, m)
	return
}
