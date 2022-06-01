package mailing

import (
	"fmt"
	"strings"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MailsConfig struct {
	MailingsConfig
	Providers map[MailProviderName]MailProviderConfigInterface `yaml:"providers"`
}

func (mmc MailsConfig) String() string {
	actions := make([]string, 0, len(mmc.Providers))

	for pName, pConfig := range mmc.Providers {
		actions = append(actions, fmt.Sprintf("%s: %s", pName.String(), pConfig.String()))
	}

	return strings.Join(actions, "; ")
}

func (mmc MailsConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(mmc)
}
