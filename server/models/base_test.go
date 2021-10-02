package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/env"
	"github.com/vavilen84/gocommerce/logger"
	"github.com/vavilen84/gocommerce/store"
	"os"
	"os/exec"
)

func beforeTestRun() {
	setTestAppEnv()
	err := godotenv.Load("./../.env.development")
	if err != nil {
		logger.LogFatal("Error loading .env file")
	}

	store.InitTestORM()
	logger.InitLogger()
}

func setTestAppEnv() {
	err := os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if err != nil {
		logger.LogError(err)
	}
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB() {
	clearTestDb()
	err := os.Chdir(os.Getenv(constants.AppRootEnvVar))
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command(
		"bee",
		"migrate",
		"refresh",
		"-driver="+env.GetSQLDriver(),
		"-conn="+env.GetDbDsn(env.GetMySQLTestDb()),
	)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	cmd = exec.Command(
		"bee",
		"migrate",
		"-driver="+env.GetSQLDriver(),
		"-conn="+env.GetDbDsn(env.GetMySQLTestDb()),
	)
	out, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	//err := orm.RunSyncdb(constants.DefaultDBAlias, true, true)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//dropAllTablesFromTestDB(ctx, conn)
	//err := CreateMigrationsTableIfNotExists(ctx, conn)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//err = MigrateUp(ctx, conn)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//LoadFixtures(ctx, conn)
	//return
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
