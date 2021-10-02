package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddProductTable_20211001_202841 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddProductTable_20211001_202841{}
	m.Created = "20211001_202841"

	migration.Register("AddProductTable_20211001_202841", m)
}

// Run the migrations
func (m *AddProductTable_20211001_202841) Up() {
	m.SQL(`
		CREATE TABLE product
		(
			id         INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title      varchar(255) NOT NULL,
			sku        varchar(255) NOT NULL,
			price      BIGINT UNSIGNED NOT NULL,
			created_at BIGINT UNSIGNED NOT NULL,
			updated_at BIGINT UNSIGNED NOT NULL,
			deleted_at BIGINT UNSIGNED NULL DEFAULT NULL
		) ENGINE=InnoDB CHARSET=utf8;
	`)
}

// Reverse the migrations
func (m *AddProductTable_20211001_202841) Down() {
	m.SQL("DROP TABLE product")
}
