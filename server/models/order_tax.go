package models

//
//import (
//	"context"
//	"orm/sql"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/orm"
//	"github.com/vavilen84/gocommerce/validator"
//	"gopkg.in/go-playground/validator.v9"
//	"log"
//)
//
//type OrderTax struct {
//	Id      uint32 `json:"id" column:"id"`
//	OrderId uint32 `json:"order_id" column:"order_id"`
//	TaxId   uint32 `json:"tax_id" column:"tax_id"`
//}
//
//func (m OrderTax) GetId() uint32 {
//	return m.Id
//}
//
//func (OrderTax) GetTableName() string {
//	return constants.OrderTaxDBTable
//}
//
//func (OrderTax) getValidationRules() validator.ScenarioRules {
//	return validator.ScenarioRules{
//		constants.ScenarioCreate: validator.FieldRules{
//			constants.OrderTaxOrderIdField: "required",
//			constants.OrderTaxTaxIdField:   "required",
//		},
//	}
//}
//
//func (OrderTax) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m OrderTax) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validator.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
