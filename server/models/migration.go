package models

import (
	"context"
	"database/sql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/helpers"
	"github.com/vavilen84/gocommerce/interfaces"
	"github.com/vavilen84/gocommerce/orm"
	"github.com/vavilen84/gocommerce/validation"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Migration struct {
	Id        uint32 `json:"id" column:"id"`
	Version   int64  `json:"version" column:"version"`
	Filename  string `json:"filename" column:"filename"`
	CreatedAt int64  `json:"created_at" column:"created_at"`
}

func (Migration) GetTableName() string {
	return constants.MigrationDBTable
}

func (m Migration) GetId() uint32 {
	return m.Id
}

func (m Migration) SetId(id uint32) interfaces.Model {
	m.Id = id
	return m
}

func (Migration) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			constants.MigrationVersionField:   "required",
			constants.MigrationFilenameField:  "required",
			constants.MigrationCreatedAtField: "required",
		},
	}
}

func (Migration) GetValidator() interface{} {
	return validator.New()
}

func getMigration(info os.FileInfo) (err error, m Migration) {
	filename := info.Name()
	splitted := strings.Split(info.Name(), "_")
	version, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Println(err)
		return
	}

	m = Migration{
		CreatedAt: time.Now().Unix(),
		Filename:  filename,
		Version:   int64(version),
	}
	return
}

func getMigrations() (err error, keys []int, list map[int64]Migration) {
	list = make(map[int64]Migration)
	keys = make([]int, 0)
	err = filepath.Walk(os.Getenv(constants.AppRootEnvVar)+"/"+constants.MigrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			helpers.LogError(err)
		}
		if !info.IsDir() {
			err, migration := getMigration(info)
			if err != nil {
				log.Println(err)
				return err
			}
			list[migration.Version] = migration
			keys = append(keys, int(migration.Version))
		}
		return nil
	})
	if err != nil {
		log.Print(err.Error())
		return
	}

	sort.Ints(keys)
	return
}

func MigrateUp(ctx context.Context, conn *sql.Conn) error {
	err, keys, list := getMigrations()
	for _, k := range keys {
		err = apply(ctx, conn, k, list)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func performMigrateTx(ctx context.Context, conn *sql.Conn, m Migration) error {
	tx, beginTxErr := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if beginTxErr != nil {
		log.Fatal(beginTxErr)
		return beginTxErr
	}

	execErr := orm.TxInsert(ctx, tx, m)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}
	content, readErr := ioutil.ReadFile(os.Getenv(constants.AppRootEnvVar) + "/" + constants.MigrationsFolder + "/" + m.Filename)
	if readErr != nil {
		log.Fatal(readErr)
		return readErr
	}
	sqlQuery := string(content)
	_, execErr = tx.ExecContext(ctx, sqlQuery)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func apply(ctx context.Context, conn *sql.Conn, k int, list map[int64]Migration) error {
	m := list[int64(k)]
	row := conn.QueryRowContext(ctx, `SELECT version FROM `+m.GetTableName()+` WHERE version = ?`, m.Version)
	var version int64
	err := row.Scan(&version)
	if err == sql.ErrNoRows {

		validationErr := validation.ValidateByScenario(constants.ScenarioCreate, &m)
		if validationErr != nil {
			log.Println(validationErr)
			return validationErr
		}

		err = performMigrateTx(ctx, conn, m)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func CreateMigrationsTableIfNotExists(ctx context.Context, conn *sql.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS ` + constants.MigrationDBTable + `
		(
   		id INT UNSIGNED NOT NULL PRIMARY KEY,
			version BIGINT UNSIGNED NOT NULL,
			filename varchar(255) NOT NULL,
			created_at BIGINT UNSIGNED NOT NULL
		) ENGINE=InnoDB CHARSET=utf8;
	`
	_, err := conn.ExecContext(ctx, query)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
