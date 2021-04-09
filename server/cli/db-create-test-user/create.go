package main

import (
	"database/sql"
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/store"
	"log"
	"os"
)

func main() {
	store.InitDB()
	conn, ctx := store.GetNewDBConn()
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf(
		`CREATE USER '%s'@'*' IDENTIFIED BY '%s';`,
		os.Getenv(constants.MysqlTestUserEnvVar),
		os.Getenv(constants.MysqlTestUserPasswordEnvVar),
	)
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf(
		`GRANT ALL PRIVILEGES ON database_name.* TO 'username'@'localhost';`,
		os.Getenv(constants.MysqlTestUserEnvVar),
		os.Getenv(constants.MysqlTestUserPasswordEnvVar),
	)
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
