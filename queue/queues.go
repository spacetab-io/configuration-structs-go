package queue

import (
	"fmt"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
	"github.com/spacetab-io/configuration-structs-go/v2/errors"
)

type ProviderName string

const (
	ProviderNSQ   ProviderName = "nsq"
	ProviderKafka ProviderName = "kafka"
)

type AvailableProvidersConfig struct {
	NSQ   NSQConfig   `yaml:"nsq"`
	Kafka KafkaConfig `yaml:"kafka"`
}

type MessageQueueConfig struct {
	Active    ProviderName             `yaml:"active" valid:"required"`
	Available AvailableProvidersConfig `yaml:"available" valid:"required"`
}

func (mqc MessageQueueConfig) String() string {
	return "message queue config"
}

func (mqc MessageQueueConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(mqc)
}

func (mqc MessageQueueConfig) GetActiveProviderConfig() (ProviderConfigInterface, error) {
	var cfg ProviderConfigInterface

	switch mqc.Active {
	case ProviderNSQ:
		cfg = &mqc.Available.NSQ
	case ProviderKafka:
		cfg = &mqc.Available.Kafka
	default:
		return nil, fmt.Errorf("%w for queueing: %s", errors.ErrUnknownProvider, mqc.Active)
	}

	return cfg, nil
}
