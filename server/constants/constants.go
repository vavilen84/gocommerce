package constants

import "time"

const (

	// db tables
	MigrationDBTable       = "migration"
	ProductDBTable         = "product"
	CustomerDBTable        = "customer"
	OrderDBTable           = "order"
	OrderProductDBTable    = "order_product"
	OrderTaxDBTable        = "order_tax"
	OrderProductTaxDBTable = "order_product_tax"
	OrderDiscountDBTable   = "order_discount"
	TaxDBTable             = "tax"

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

	//tax
	TaxTitleField      = "Title"
	TaxAmountField     = "Amount"
	TaxPercentageField = "Percentage"
	TaxTypeField       = "Type"

	//tax types
	TaxCartType     = 1
	TaxCategoryType = 2
	TaxProductType  = 3

	//discount
	DiscountTitleField      = "Title"
	DiscountAmountField     = "Amount"
	DiscountPercentageField = "Percentage"
	DiscountTypeField       = "Type"

	//discount types
	DiscountCartType     = 1
	DiscountCategoryType = 2
	DiscountProductType  = 3

	//order
	OrderCustomerIdField = "CustomerId"

	//order_product
	OrderOrderIdField   = "OrderId"
	OrderProductIdField = "ProductId"
	OrderQuantityField  = "Quantity"

	//order_tax
	OrderTaxOrderIdField = "OrderId"
	OrderTaxTaxIdField   = "TaxId"

	//order_discount
	OrderDiscountOrderIdField    = "OrderId"
	OrderDiscountDiscountIdField = "DiscountId"

	//order_tax
	OrderProductTaxOrderProductIdField = "OrderProductId"
	OrderProductTaxTaxIdField          = "TaxId"
)
