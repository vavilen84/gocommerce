package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/gocommerce/constants"
	"os"
)

var (
	testDb *sql.DB
)

func InitTestDB() {
	testDb = initTestDb()
}

func GetTestDB() *sql.DB {
	return testDb
}

func initTestDb() *sql.DB {
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
