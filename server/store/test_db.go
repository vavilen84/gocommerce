package store

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

func initTestDBForLocalhostAppRun() *sql.DB {
	// TODO dirty hack, needs refactoring - main_test.go and models/* tests have different relative path to  .env
	err := godotenv.Load("../../.env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			helpers.LogError(err)
			os.Exit(1)
		}
	}
	sqlServerDsn := os.Getenv("LOCALHOST_SQL_DSN")
	mysqlDbName := os.Getenv("MYSQL_TEST_DATABASE")
	DbDsn := os.Getenv("LOCALHOST_TEST_DB_SQL_DSN")
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}

func initTestDBForDockerMySql() *sql.DB {
	sqlServerDsn := os.Getenv("SQL_DSN")
	mysqlDbName := os.Getenv("MYSQL_TEST_DATABASE")
	DbDsn := os.Getenv("TEST_DB_SQL_DSN")
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}

func initTestDBForHostMachineMySql() *sql.DB {
	sqlServerDsn := setHostAddress(os.Getenv("HOST_MACHINE_SQL_DSN"))
	mysqlDbName := setHostAddress(os.Getenv("MYSQL_TEST_DATABASE"))
	DbDsn := setHostAddress(os.Getenv("HOST_MACHINE_TEST_SQL_DSN"))
	return processInitDb(sqlServerDsn, mysqlDbName, DbDsn)
}

func initTestDb() *sql.DB {
	docker := os.Getenv("DOCKER")
	if docker != "true" {
		return initTestDBForLocalhostAppRun()
	}
	dockerMySqlOnHostMachine := os.Getenv("DOCKER_MYSQL_HOST_MACHINE")
	if dockerMySqlOnHostMachine != "true" {
		return initTestDBForDockerMySql()
	}
	return initTestDBForHostMachineMySql()
}
