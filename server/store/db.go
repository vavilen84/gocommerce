package store

import (
	"database/sql"
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"os"
)

var (
	db *sql.DB
)

func InitDB() {
	db = initDb()
}

func GetDB() *sql.DB {
	// for endpoints testing. usage: os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if os.Getenv(constants.AppEnvEnvVar) == constants.TestingAppEnv {
		return GetTestDB()
	}
	return db
}

func initDb() *sql.DB {
	sqlServerDsn := fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		"",
	)
	mysqlDbName := os.Getenv(constants.MysqlDBEnvVar)
	DbDsn := fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		os.Getenv(constants.MysqlDBEnvVar),
	)
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}
