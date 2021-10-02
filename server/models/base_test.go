package models

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/env"
	"github.com/vavilen84/gocommerce/helpers"
	"github.com/vavilen84/gocommerce/logger"
	"github.com/vavilen84/gocommerce/store"
	"os"
	"path"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	var once sync.Once
	onceCall := func() {
		beforeAllTestRun()
	}
	once.Do(onceCall)
	code := m.Run()
	os.Exit(code)
}

func beforeEachTest() {
	restoreFromDump()
}

func beforeAllTestRun() {
	setTestAppEnv()
	store.InitTestORM()
	logger.InitLogger()
	prepareTestDB()
}

func setTestAppEnv() {
	err := os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if err != nil {
		logger.LogError(err)
	}
	// allow to run tests from app root & from packages folders
	err = godotenv.Load(".env.development")
	if err != nil {
		err = godotenv.Load("../.env.development")
		if err != nil {
			logger.LogFatal("Error loading .env file")
		}
	}
}

func runMigrations() {
	helpers.RunCmd(
		"bee",
		"migrate",
		"-driver="+env.GetSQLDriver(),
		"-conn="+env.GetDbDsn(env.GetMySQLTestDb()),
	)
}

func restoreFromDump() {
	helpers.RunCmd(
		"mysql",
		"-u"+env.GetMySQLUser(),
		"-p"+env.GetMySQLUserPass(),
		env.GetMySQLTestDb(),
		"<",
		getDbDumpFile(),
	)
}

func getDbDumpFile() string {
	return path.Join(env.GetAppRoot(), constants.TmpFolder, constants.TestDbDumpFile)
}

func createDbDump() {
	os.Remove(getDbDumpFile())
	helpers.RunCmd(
		"mysqldump",
		"-u"+env.GetMySQLUser(),
		"-p"+env.GetMySQLUserPass(),
		env.GetMySQLTestDb(),
		"--result-file="+getDbDumpFile(),
	)
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB() {
	clearTestDb()
	runMigrations()
	createDbDump()
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func clearTestDb() {
	// use credentials without db in order to create db
	db, err := sql.Open(env.GetSQLDriver(), env.GetSQLServerDsn())
	if err != nil {
		logger.LogFatal(err)
	}
	ctx := store.GetDefaultDBContext()
	conn, err := db.Conn(ctx)
	if err != nil {
		logger.LogFatal(err)
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, "DROP DATABASE "+env.GetMySQLTestDb())
	if err != nil {
		logger.LogFatal(err)
	}
	_, err = conn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+env.GetMySQLTestDb())
	if err != nil {
		logger.LogFatal(err)
	}
}
