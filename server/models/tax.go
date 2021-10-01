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
//type Tax struct {
//	Id         uint32 `json:"id" column:"id"`
//	Title      string `json:"title" column:"title"`
//	Amount     uint32 `json:"amount" column:"amount"`
//	Percentage uint8  `json:"percentage" column:"percentage"`
//	Type       uint8  `json:"type" column:"type"`
//
//	CreatedAt int64 `json:"created_at" column:"created_at"`
//	UpdatedAt int64 `json:"updated_at" column:"updated_at"`
//	DeletedAt int64 `json:"deleted_at" column:"deleted_at"`
//}
//
//func (m Tax) GetId() uint32 {
//	return m.Id
//}
//
//func (Tax) GetTableName() string {
//	return constants.TaxDBTable
//}
//
//func (Tax) getValidationRules() validator.ScenarioRules {
//	return validator.ScenarioRules{
//		constants.ScenarioCreate: validator.FieldRules{
//			constants.TaxTitleField:      "required",
//			constants.TaxAmountField:     "min=1",
//			constants.TaxPercentageField: "min=1,max=100",
//			constants.TaxTypeField:       "required",
//		},
//	}
//}
//
//func (Tax) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m Tax) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validator.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
