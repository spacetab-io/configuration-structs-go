package queue

import (
	"fmt"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type ProviderConfigInterface interface {
	contracts.AdressesInterface

	GetLogLevel() string
}

type ConfigInterface interface {
	contracts.ValidatableInterface
	fmt.Stringer

	GetActiveProviderConfig() (ProviderConfigInterface, error)
}
