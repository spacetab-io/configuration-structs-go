package cfgstructs

import (
	"fmt"
)

type NSQQueue struct {
	Enable      bool    `yaml:"enable"`
	LogLevel    string  `yaml:"log_level"`
	LookupdHost string  `yaml:"lookupd_host"`
	LookupdPort int     `yaml:"lookupd_port"`
	NsqdPort    int     `yaml:"nsqd_port"`
	NsqdHost    string  `yaml:"nsqd_host"`
	MaxInFlight int     `yaml:"max_in_flight"`
	MaxAttempts *uint16 `yaml:"max_attempts"`
}

type MessageQueue struct {
	Nsq NSQQueue `yaml:"nsq"`
}

func (s NSQQueue) GetNSQLookupdPath() string {
	return fmt.Sprintf("%s:%v", s.LookupdHost, s.LookupdPort)
}

func (s NSQQueue) GetNSQdPath() string {
	return fmt.Sprintf("%s:%v", s.NsqdHost, s.NsqdPort)
}
