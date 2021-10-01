package store

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
	"os"
)

func InitTestORM() {
	registerDriver()
	registerDatabase(os.Getenv(constants.MysqlTestDBEnvVar))
	orm.Debug = true
}

func InitORM() {
	registerDriver()
	registerDatabase(os.Getenv(constants.MysqlDBEnvVar))
}

func registerDriver() {
	err := orm.RegisterDriver(os.Getenv(constants.SqlDriverEnvVar), orm.DRMySQL)
	if err != nil {
		logger.LogFatal(err)
	}
}

func registerDatabase(mysqlDbName string) {
	dbDsn := fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		mysqlDbName,
	)
	err := orm.RegisterDataBase(constants.DefaultDBAlias, os.Getenv(constants.SqlDriverEnvVar), dbDsn, 10, 10)
	if err != nil {
		logger.LogFatal(err)
	}
}
