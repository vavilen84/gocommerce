package main

import (
	"fmt"
	"github.com/vavilen84/gocommerce/store"
	"log"
	"os"
)

func main() {
	store.InitDB()
	conn, ctx := store.GetNewDBConn()
	defer conn.Close()

	query := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS %s",
		os.Getenv("MYSQL_TEST_DATABASE"),
	)

	_, err := conn.ExecContext(ctx, query)
	if err != nil {
		log.Print(err.Error())
	}
}
