package cfgstructs

import (
	"fmt"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type SeedInfo struct {
	Enabled     bool   `yaml:"enabled" valid:"required"`
	Name        string `yaml:"name" valid:"optional"`
	Description string `yaml:"description" valid:"optional"`
	ClassName   string `yaml:"class_name" valid:"optional"`
}

type SeedingCfg struct {
	RunOnStart bool       `yaml:"run_on_start" valid:"type(bool)"`
	Seeds      []SeedInfo `yaml:"seeds" valid:"optional"`
}

type MigrationCfg struct {
	RunOnStart bool   `yaml:"run_on_start" valid:"type(bool)"`
	TableName  string `yaml:"table_name" valid:"required"`
	Path       string `yaml:"path" valid:"required"`
}

type Database struct {
	Driver             string             `yaml:"driver" valid:"required,in(postgres|mysql)"`
	Hosts              contracts.HostsCfg `yaml:"hosts" valid:"required"`
	User               string             `yaml:"user" valid:"required"`
	Pass               string             `yaml:"pass" valid:"optional"`
	Schema             string             `yaml:"schema" valid:"optional"`
	Name               string             `yaml:"database" valid:"required"`
	SSLMode            string             `yaml:"ssl_mode" valid:"optional,in(disable|enable)"`
	LogLevel           string             `yaml:"log_level" valid:"required,in(trace|debug|info|warning|error)"`
	Seeding            SeedingCfg         `yaml:"seeding" valid:"optional"`
	Migrations         MigrationCfg       `yaml:"migrations" valid:"optional"`
	MaxOpenConnections int                `yaml:"max_open_connections" valid:"optional,int"`
	MaxIdleConnections int                `yaml:"max_idle_connections" valid:"optional,int"`
	ConnectionLifeTime time.Duration      `yaml:"connection_lifetime" valid:"optional,duration"`
}

func (d Database) String() string {
	return "database"
}

func (d Database) GetConnectionURL() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s/%s?search_path=%s&sslmode=%s",
		d.Driver,
		d.User,
		d.Pass,
		strings.Join(d.Hosts.GetHostPortPairs(), ","),
		d.Name,
		d.Schema,
		d.SSLMode,
	)
}

func (d Database) GetDSN() string {
	if d.Driver != "postgres" {
		return d.GetConnectionURL()
	}

	return fmt.Sprintf(
		"host=%s port=%s search_path=%s dbname=%s user=%s password=%s sslmode=%s",
		strings.Join(d.Hosts.GetHosts(), ","),
		strings.Join(d.Hosts.GetPorts(), ","),
		d.Schema,
		d.Name,
		d.User,
		d.Pass,
		d.SSLMode,
	)
}

func (d Database) GetMigrationDSN() string {
	return fmt.Sprintf("%s&x-migrations-table=%s", d.GetDSN(), d.Migrations.TableName)
}

func (d Database) GetMigrationsPath() string {
	return d.Migrations.Path
}

func (d Database) GetMigrationsTableName() string {
	return d.Migrations.TableName
}

func (d Database) GetSchema() string {
	return d.Schema
}

func (d Database) GetConnectionParams() (maxConnLifetime time.Duration, maxConns, minConns int32) {
	return d.ConnectionLifeTime, int32(d.MaxOpenConnections), int32(d.MaxIdleConnections)
}

func (d Database) MigrateOnStart() bool {
	return d.Migrations.RunOnStart
}

func (d Database) SeedOnStart() bool {
	return d.Seeding.RunOnStart
}

func (d Database) Validate() (bool, error) {
	govalidator.CustomTypeTagMap.Set("duration", func(dur interface{}, o interface{}) bool {
		var err error

		switch v := dur.(type) {
		case string:
			_, err = time.ParseDuration(v)
		case time.Duration:
			err = nil
		}

		return err == nil
	})

	return contracts.ConfigValidate(d)
}
