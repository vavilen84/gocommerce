package constants

import "time"

const (

	// struct names

	ProductStructName = "Product"

	// db table names

	ProductTableName    = "product"
	MigrationsTableName = "migrations"

	// migrations

	MigrationsFolder = "migrations"

	// scenarios

	ScenarioCreate = "create"
	ScenarioUpdate = "update"

	// validation tags

	RequiredTag = "required"
	MinTag      = "min"
	MaxTag      = "max"
	EmailTag    = "email"

	// validation error messages

	RequiredErrorMsg = "%s resource: '%s' is required"
	MinValueErrorMsg = "%s resource: '%s' min value is %s"
	MaxValueErrorMsg = "%s resource: '%s' max value is %s"
	Uuid4ErrorMsg    = "%s resource: '%s' is not valid uuid4"
	EmailErrorMsg    = "%s resource: email is not valid"

	// Server

	DefaultWriteTimout  = 60 * time.Second
	DefaultReadTimeout  = 60 * time.Second
	DefaultStoreTimeout = 60 * time.Second
)
