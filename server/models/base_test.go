package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/joho/godotenv"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
	"github.com/vavilen84/gocommerce/store"
	"log"
	"os"
)

func beforeTestRun() {
	setTestAppEnv()
	err := godotenv.Load("./../.env.development")
	if err != nil {
		logger.LogFatal("Error loading .env file")
	}

	store.InitTestORM()
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

	err := orm.RunSyncdb(constants.DefaultDBAlias, true, true)
	if err != nil {
		fmt.Println(err)
	}

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
func dropAllTablesFromTestDB(ctx context.Context, conn *sql.Conn) {
	tables := []string{
		constants.MigrationDBTable,
		constants.ProductDBTable,
	}
	for i := 0; i < len(tables); i++ {
		_, err := conn.ExecContext(ctx, "DROP TABLE IF EXISTS "+tables[i])
		if err != nil {
			log.Println(err)
		}
	}
}
