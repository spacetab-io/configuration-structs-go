package cfgstructs

import (
	"time"
)

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
	Summary() string
}

type DatabaseCfgInterface interface {
	GetConnectionURL() string
	GetDSN() string
	GetMigrationDSN() string
	GetMigrationsPath() string

	GetSchema() string
	GetMigrationsTableName() string

	GetConnectionParams() (maxConnLifetime time.Duration, maxConns, minConns int32)
}

type NSQQueueCfgInterface interface {
	GetNSQLookupdPath() string
	GetNSQdPath() string
}
