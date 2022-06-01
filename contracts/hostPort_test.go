package contracts_test

import (
	"testing"

	cfgstructs "github.com/spacetab-io/configuration-structs-go/v2/contracts"
	"github.com/stretchr/testify/assert"
)

func TestHostCfg_String(t *testing.T) {
	type testCase struct {
		name string
		in   cfgstructs.HostCfg
		exp  string
	}

	tcs := []testCase{
		{
			name: "random host-port pair",
			in:   cfgstructs.HostCfg{Host: "some.host.dm", Port: 432},
			exp:  "some.host.dm:432",
		},
		{
			name: "https host-port pair",
			in:   cfgstructs.HostCfg{Host: "https://some.host.dm", Port: 443},
			exp:  "https://some.host.dm",
		},
		{
			name: "http host-port pair",
			in:   cfgstructs.HostCfg{Host: "http://some.host.dm", Port: 80},
			exp:  "http://some.host.dm",
		},
		{
			name: "ssl host-port pair",
			in:   cfgstructs.HostCfg{Host: "ssl://some.host.dm", Port: 21},
			exp:  "ssl://some.host.dm",
		},
		{
			name: "empty port http",
			in:   cfgstructs.HostCfg{Host: "http://some.host.dm", Port: 0},
			exp:  "http://some.host.dm",
		},
		{
			name: "empty port random host-port pair",
			in:   cfgstructs.HostCfg{Host: "some.host.dm", Port: 0},
			exp:  "some.host.dm",
		},
	}

	t.Parallel()

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.exp, tc.in.String())
		})
	}
}
