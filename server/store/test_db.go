package store

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/helpers"
	"os"
)

var (
	testDb *sql.DB
)

func InitTestDB() {
	testDb = initTestDb()
}

func GetNewTestDBConn() (conn *sql.Conn, ctx context.Context) {
	ctx = GetDefaultDBContext()
	conn, err := testDb.Conn(ctx)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func initTestDb() *sql.DB {
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
		os.Getenv(constants.MysqlTestDBEnvVar),
	)
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}
