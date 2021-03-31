package models

import (
	"context"
	"database/sql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/store"
	"log"
)

func init() {
	store.InitTestDB()
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB(ctx context.Context, conn *sql.Conn) {
	dropAllTablesFromTestDB(ctx, conn)
	err := CreateMigrationsTableIfNotExists(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	err = MigrateUp(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	LoadFixtures(ctx, conn)
	return
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
