package cfgstructs_test

import (
	"testing"
	"time"

	cfgstructs "github.com/spacetab-io/configuration-structs-go/v2"
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
	"github.com/stretchr/testify/assert"
)

func TestDatabase_Validate(t *testing.T) {
	type testCase struct {
		name     string
		in       contracts.ValidatableInterface
		hasError bool
	}

	defaultDatabaseCfg := dbConfig(t)
	emptyDatabaseCfg := cfgstructs.Database{}

	tcs := []testCase{
		{
			name: "no errors",
			in:   defaultDatabaseCfg,
		},
		{
			name:     "fail params",
			in:       emptyDatabaseCfg,
			hasError: true,
		},
	}

	t.Parallel()

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := tc.in.Validate()
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func dbConfig(t *testing.T) cfgstructs.Database {
	t.Helper()

	return cfgstructs.Database{
		Driver:             "postgres",
		Hosts:              contracts.HostsCfg{{Host: "127.0.0.1", Port: 5432}},
		User:               "postgres",
		Pass:               "",
		Schema:             "public",
		Name:               "db_name",
		SSLMode:            "disable",
		LogLevel:           "warning",
		Seeding:            cfgstructs.SeedingCfg{RunOnStart: false, Seeds: nil},
		Migrations:         cfgstructs.MigrationCfg{RunOnStart: false, TableName: "migrations", Path: "./"},
		MaxOpenConnections: 10,
		MaxIdleConnections: 0,
		ConnectionLifeTime: 30 * time.Second,
	}
}
