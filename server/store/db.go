package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/helpers"
	"os"
)

var (
	db *sql.DB
)

func InitDB() {
	db = initDb()
}

func GetNewDBConn() (conn *sql.Conn, ctx context.Context) {
	// for endpoints testing. usage: os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if os.Getenv(constants.AppEnvEnvVar) == constants.TestingAppEnv {
		return GetNewTestDBConn()
	}
	ctx = GetDefaultDBContext()
	conn, err := db.Conn(ctx)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func initDb() *sql.DB {
	sqlServerDsn := fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.DockerMysqlServiceEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		"",
	)
	mysqlDbName := os.Getenv(constants.MysqlDBEnvVar)
	DbDsn := fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.DockerMysqlServiceEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		os.Getenv(constants.MysqlDBEnvVar),
	)
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}
