package models

import (
	"context"
	"database/sql"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/database"
	"github.com/vavilen84/gocommerce/helpers"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Migration struct {
	BaseModel
	Id       uint32 `json:"id" column:"id"`
	Version  int64  `json:"version" column:"version"`
	Filename string `json:"filename" column:"filename"`
}

func (Migration) GetTableName() string {
	return constants.MigrationsTableName
}

func (m *Migration) ValidateByScenario() {
	validationMap := make(ValidationMap)
	validationMap = ValidationMap{
		constants.ScenarioCreate: ValidationRules{
			"CreatedAt": "required",
			"UpdatedAt": "required",
			"Version":   "required",
			"Filename":  "required",
		},
	}
	m.ValidationErrors = validateByScenario(m, validationMap)
}

func (m Migration) GetId() uint32 {
	return m.Id
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
		Filename: filename,
		Version:  int64(version),
	}
	return
}

func getMigrations() (err error, keys []int, list map[int64]Migration) {
	list = make(map[int64]Migration)
	keys = make([]int, 0)
	err = filepath.Walk(os.Getenv("PROJECT_ROOT")+"/"+constants.MigrationsFolder, func(path string, info os.FileInfo, err error) error {
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

	execErr := database.TxInsert(ctx, tx, &m)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}
	content, readErr := ioutil.ReadFile(os.Getenv("PROJECT_ROOT") + "/" + constants.MigrationsFolder + "/" + m.Filename)
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
	row := conn.QueryRowContext(ctx, `SELECT version FROM `+constants.MigrationsTableName+` WHERE version = ?`, m.Version)
	var version int64
	err := row.Scan(&version)
	if err == sql.ErrNoRows {

		m.Scenario = constants.ScenarioCreate
		m.ValidateByScenario()
		if len(m.ValidationErrors) > 0 {
			log.Println(m.ValidationErrors)
			return helpers.MergeErrors(m.ValidationErrors)
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
		CREATE TABLE IF NOT EXISTS ` + constants.MigrationsTableName + `
		(
    		id INT UNSIGNED NOT NULL PRIMARY KEY,
			version BIGINT UNSIGNED NOT NULL,
			filename varchar(255) NOT NULL,
			created_at BIGINT UNSIGNED NOT NULL,
			updated_at BIGINT UNSIGNED NOT NULL
		) ENGINE=InnoDB CHARSET=utf8;
	`
	_, err := conn.ExecContext(ctx, query)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func (m *Migration) Validate() {
	m.ValidationErrors = validate(m)
}
