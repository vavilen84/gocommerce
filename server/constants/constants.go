package constants

import "time"

const (

	// db tables
	MigrationDBTable = "migration"
	ProductDBTable   = "product"
	CustomerDBTable  = "customer"
	OrderDBTable     = "order"

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
	EmailErrorMsg    = "%s resource: email is not valid"

	// Server
	DefaultStoreTimeout = 60 * time.Second

	// model names
	MigrationModel = "Migration"

	// field names

	// migration
	MigrationVersionField   = "Version"
	MigrationFilenameField  = "Filename"
	MigrationCreatedAtField = "CreatedAt"
	MigrationUpdatedAtField = "UpdatedAt"

	//product
	ProductTitleField = "Title"
	ProductSKUField   = "SKU"
	ProductPriceField = "Price"

	//customer
	CustomerFirstNameField = "FirstName"
	CustomerLastNameField  = "LastName"
	CustomerEmailField     = "Email"

	//order
	OrderCustomerIdField = "CustomerId"
)
