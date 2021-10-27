package cfgstructs

type Config interface {
	ValidateAll() []error
}

type Configurator interface {
	Validate() []error
}

type ApplicationInfoCfgInterface interface {
	GetString() string
	GetAlias() string
	GetVersion() string
}

type DatabaseCfgInterface interface {
	GetConnectionURL() string
	GetDSN() string
	GetMigrationDSN() string
	GetMigrationsPath() string

	GetSchema() string
	GetMigrationsTableName() string
}

type NSQQueueCfgInterface interface {
	GetNSQLookupdPath() string
	GetNSQdPath() string
}
