package cfgstructs

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	connMaxLifeTime          = 30 * time.Minute
	maxOpenConn              = 4
	defaultSchema            = "public"
	cfgDBPrefix              = "[config/database]"
	cfgValidationErrorFormat = "%s: %w - %s"
)

type SeedingCfg struct {
	RunOnStart bool     `yaml:"run_on_start"`
	Seeds      []string `yaml:"seeds"`
}

type MigrationCfg struct {
	RunOnStart bool   `yaml:"run_on_start"`
	TableName  string `yaml:"table_name"`
	Path       string `yaml:"path"`
}

type Database struct {
	Driver             string        `yaml:"driver"`
	Host               string        `yaml:"host"`
	Port               uint          `yaml:"port"`
	User               string        `yaml:"user"`
	Pass               string        `yaml:"pass"`
	Schema             string        `yaml:"schema"`
	Name               string        `yaml:"database"`
	SSLMode            string        `yaml:"ssl_mode"`
	LogLevel           string        `yaml:"log_level"`
	Seeding            SeedingCfg    `yaml:"seeding"`
	Migrations         MigrationCfg  `yaml:"migrations"`
	MaxOpenConnections int           `yaml:"max_open_connections"`
	MaxIdleConnections int           `yaml:"max_idle_connections"`
	ConnectionLifeTime time.Duration `yaml:"connection_lifetime"`
}

func (d *Database) GetConnectionURL() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?search_path=%s&sslmode=%s",
		d.Driver,
		d.User,
		d.Pass,
		d.Host,
		d.Port,
		d.Name,
		d.Schema,
		d.SSLMode,
	)
}

func (d *Database) GetDSN() string {
	if d.Driver != "postgres" {
		return d.GetConnectionURL()
	}

	return fmt.Sprintf(
		"host=%s port=%d search_path=%s dbname=%s user=%s password=%s sslmode=%s",
		d.Host,
		d.Port,
		d.Schema,
		d.Name,
		d.User,
		d.Pass,
		d.SSLMode,
	)
}

func (d *Database) GetMigrationDSN() string {
	return fmt.Sprintf("%s&x-migrations-table=migrations", d.GetDSN())
}

func (d *Database) GetMigrationsPath() string {
	return fmt.Sprintf("file://%s", d.Migrations.Path)
}

func (d Database) GetMigrationsTableName() string {
	return d.Migrations.TableName
}

func (d Database) GetSchema() string {
	return d.Schema
}

var (
	ErrDataIsEmpty   = errors.New("data is empty")
	ErrDataIsInvalid = errors.New("data is invalid")
)

func (d *Database) Validate() []error {
	errList := make([]error, 0)

	d.Driver = strings.TrimSpace(d.Driver)
	switch d.Driver {
	case "postgres", "mysql":
		break
	case "":
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, ErrDataIsEmpty, "driver"))
	default:
		errList = append(errList, fmt.Errorf(
			cfgValidationErrorFormat,
			cfgDBPrefix,
			ErrDataIsEmpty,
			"driver is unknown. only 'postgres' and 'mysql' are well-known",
		))
	}

	d.User = strings.TrimSpace(d.User)
	if d.User == "" {
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, ErrDataIsEmpty, "user"))
	}

	// d.Password = strings.TrimSpace(d.Password)
	// if d.Password == "" {
	// 	errList = append(errList, fmt.Errorf(cfgDBPrefix, "password is empty"))
	// }

	d.Host = strings.TrimSpace(d.Host)
	if d.Host == "" {
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, ErrDataIsEmpty, "host"))
	} else if _, err := net.LookupHost(d.Host); err != nil {
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, err, "host resolve error"))
	}

	if d.Port == 0 || d.Port > 70000 {
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, ErrDataIsInvalid, "port"))
	}

	d.Schema = strings.TrimSpace(d.Schema)
	if d.Schema == "" {
		d.Schema = defaultSchema
	}

	d.Name = strings.TrimSpace(d.Name)
	if d.Name == "" {
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, ErrDataIsEmpty, "db name"))
	}

	if d.ConnectionLifeTime == 0 {
		d.ConnectionLifeTime = connMaxLifeTime
	}

	if d.MaxOpenConnections == 0 {
		d.MaxOpenConnections = maxOpenConn
	}

	d.Migrations.Path = strings.TrimSpace(d.Migrations.Path)
	if len(d.Migrations.Path) == 0 {
		errList = append(errList, fmt.Errorf(cfgValidationErrorFormat, cfgDBPrefix, ErrDataIsEmpty, "migrations_path"))
	}

	return errList
}
