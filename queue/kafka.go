package queue

import (
	"fmt"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type KafkaConfig struct {
	LogLevel string              `yaml:"log_level" valid:"required"`
	Brokers  []contracts.HostCfg `yaml:"brokers" valid:"required"`
}

func (kq KafkaConfig) GetHosts() []string {
	hh := make([]string, 0, len(kq.Brokers))

	for _, b := range kq.Brokers {
		hh = append(hh, b.GetHost())
	}

	return hh
}

func (kq KafkaConfig) GetPorts() []uint {
	pp := make([]uint, 0, len(kq.Brokers))

	for _, b := range kq.Brokers {
		pp = append(pp, b.GetPort())
	}

	return pp
}

func (kq KafkaConfig) GetHostPortPairs() map[string]string {
	hpp := make(map[string]string)

	for i, b := range kq.Brokers {
		hpp[fmt.Sprint(i)] = b.String()
	}

	return hpp
}

func (kq KafkaConfig) GetLogLevel() string {
	return kq.LogLevel
}
