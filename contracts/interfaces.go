package contracts

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type Config interface {
	ValidateAll() []error
}

type ValidatableInterface interface {
	fmt.Stringer

	Validate() (bool, error)
}

type ConfigValidatorInterface interface {
	Add(cfg ValidatableInterface)
	GetConfigs() []ValidatableInterface
	Validate() error
}

type AddressInterface interface {
	fmt.Stringer

	IsEmpty() bool
	GetHost() string
	GetPort() uint
}

type AdressesInterface interface {
	GetHosts() []string
	GetPorts() []uint
	GetHostPortPairs() map[string]string
}

type ApplicationInfoCfgInterface interface {
	GetString() string
	GetAlias() string
	GetVersion() string
	Summary() string
}

type LogsCfgInterface interface {
	GetLevel() string
	GetFormat() string
	SetFormat(format string)
	IsColored() bool
	GetSentryParams() (enable bool, debug bool, dsn string)
	ShowCaller() bool
	GetCallerSkipFrames() int
	SetCaller(isDisabled bool, skipFrames int)
	IsSentryEnabled() bool
	SentryDebugEnabled() bool
	GetSentryDSN() string
}

type DatabaseCfgInterface interface {
	GetConnectionURL() string
	GetDSN() string
	GetMigrationDSN() string
	GetMigrationsPath() string

	GetSchema() string
	GetMigrationsTableName() string
	MigrateOnStart() bool
	SeedOnStart() bool
	GetConnectionParams() (maxConnLifetime time.Duration, maxConns, minConns int32)
}

type SideRestServiceInterface interface {
	DebugEnable() bool
	GzipContent() bool
	GetTimeout() time.Duration
	GetBaseURL() string
}

type WebServerInterface interface {
	ValidatableInterface

	GetReadRequestTimeout() time.Duration
	GetWriteResponseTimeout() time.Duration
	GetIdleTimeout() time.Duration
	GetShutdownTimeout() time.Duration
	GetMaxConnsPerIP() int
	GetMaxRequestsPerConn() int
	UseCompression() bool
	CORSEnabled() bool
	GetListenAddress() string
}

func ConfigValidate(mp ValidatableInterface) (bool, error) {
	ok, err := govalidator.ValidateStruct(mp)
	if err != nil {
		return false, fmt.Errorf("%s config validate error: %w", mp.String(), err)
	}

	return ok, nil
}
