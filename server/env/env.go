package env

import (
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"os"
)

func GetDbDsn(dbname string) string {
	return fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		dbname,
	)
}

func GetSQLServerDsn() string {
	return fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		"",
	)
}

func GetMySQLUser() string {
	return os.Getenv(constants.MysqlUserEnvVar)
}

func GetMySQLUserPass() string {
	return os.Getenv(constants.MysqlPasswordEnvVar)
}

func GetSQLDriver() string {
	return os.Getenv(constants.SqlDriverEnvVar)
}

func GetMySQLDb() string {
	return os.Getenv(constants.MysqlDBEnvVar)
}

func GetMySQLTestDb() string {
	return os.Getenv(constants.MysqlTestDBEnvVar)
}

func GetAppRoot() string {
	return os.Getenv(constants.AppRootEnvVar)
}
