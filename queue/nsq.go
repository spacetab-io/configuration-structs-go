package queue

import (
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type NSQConfig struct {
	LogLevel    string            `yaml:"log_level" valid:"required"`
	Lookupd     contracts.HostCfg `yaml:"lookupd" valid:"required"`
	Nsqd        contracts.HostCfg `yaml:"nsqd" valid:"required"`
	MaxInFlight int               `yaml:"max_in_flight" valid:"required"`
	MaxAttempts *uint16           `yaml:"max_attempts" valid:"required"`
}

func (nq NSQConfig) GetHosts() []string {
	hh := make([]string, 0, 2) // nolint: gomnd
	hh = append(hh, nq.Lookupd.GetHost())
	hh = append(hh, nq.Nsqd.GetHost())

	return hh
}

func (nq NSQConfig) GetPorts() []uint {
	pp := make([]uint, 0, 2) // nolint: gomnd
	pp = append(pp, nq.Lookupd.GetPort())
	pp = append(pp, nq.Nsqd.GetPort())

	return pp
}

func (nq NSQConfig) GetHostPortPairs() map[string]string {
	hpp := make(map[string]string)

	hpp["lookupd"] = nq.Lookupd.String()
	hpp["nsqd"] = nq.Nsqd.String()

	return hpp
}

func (nq NSQConfig) GetLogLevel() string {
	return nq.LogLevel
}
