package store

import (
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/env"
	"github.com/vavilen84/gocommerce/logger"
)

func InitTestORM() {
	registerDriver()
	registerDatabase(env.GetMySQLTestDb())
	orm.Debug = true
}

func InitORM() {
	registerDriver()
	registerDatabase(env.GetMySQLDb())
}

func registerDriver() {
	err := orm.RegisterDriver(env.GetSQLDriver(), orm.DRMySQL)
	if err != nil {
		logger.LogFatal(err)
	}
}

func registerDatabase(mysqlDbName string) {
	err := orm.RegisterDataBase(
		constants.DefaultDBAlias,
		env.GetSQLDriver(),
		env.GetDbDsn(mysqlDbName),
		10,
		10,
	)
	if err != nil {
		logger.LogFatal(err)
	}
}
