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
//type OrderProduct struct {
//	Id        uint32 `json:"id" column:"id"`
//	OrderId   uint32 `json:"order_id" column:"order_id"`
//	ProductId uint32 `json:"product_id" column:"product_id"`
//	Quantity  uint8  `json:"quantity" column:"quantity"`
//}
//
//func (m OrderProduct) GetId() uint32 {
//	return m.Id
//}
//
//func (OrderProduct) GetTableName() string {
//	return constants.OrderProductDBTable
//}
//
//func (OrderProduct) getValidationRules() validator.ScenarioRules {
//	return validator.ScenarioRules{
//		constants.ScenarioCreate: validator.FieldRules{
//			constants.OrderOrderIdField:   "required",
//			constants.OrderProductIdField: "required",
//			constants.OrderQuantityField:  "required,min=1,max=255",
//		},
//	}
//}
//
//func (OrderProduct) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m OrderProduct) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validator.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
