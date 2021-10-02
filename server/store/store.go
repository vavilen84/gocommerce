package store

import (
	"context"
	"database/sql"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/env"
	"github.com/vavilen84/gocommerce/logger"
)

func InitTestORM() {
	registerDriver()
	createDbIfNotExists(env.GetMySQLTestDb())
	registerDatabase(env.GetMySQLTestDb())
	orm.Debug = true
}

func InitORM() {
	registerDriver()
	createDbIfNotExists(env.GetMySQLDb())
	registerDatabase(env.GetMySQLDb())
}

func GetDefaultDBContext() context.Context {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, constants.DefaultStoreTimeout)
	return ctx
}

func createDbIfNotExists(dbName string) {
	// use credentials without db in order to create db
	db, err := sql.Open(env.GetSQLDriver(), env.GetSQLServerDsn())
	if err != nil {
		logger.LogFatal(err)
	}
	ctx := GetDefaultDBContext()
	conn, err := db.Conn(ctx)
	if err != nil {
		logger.LogFatal(err)
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		logger.LogFatal(err)
	}
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
