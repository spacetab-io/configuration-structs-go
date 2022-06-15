package mailing

import (
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MandrillConfig struct {
	Key   string `yaml:"key" valid:"required"`
	Async bool   `yaml:"isAsync" valid:"-"`
}

func (mc MandrillConfig) String() string {
	return mc.Name().String()
}

func (mc MandrillConfig) Name() MailProviderName {
	return MailProviderMandrill
}

func (mc MandrillConfig) GetHostPort() contracts.AddressInterface {
	return nil
}

func (mc MandrillConfig) GetEncryption() MailProviderEncryption {
	return MailProviderEncryptionNone
}

func (mc MandrillConfig) GetAuthType() contracts.AuthType {
	return contracts.AuthTypeNone
}

func (mc MandrillConfig) GetDKIMPrivateKey() *string {
	return nil
}

func (mc MandrillConfig) GetConnectionTimeout() time.Duration {
	return 0
}

func (mc MandrillConfig) GetSendTimeout() time.Duration {
	return 0
}

func (mc MandrillConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(mc)
}

func (mc MandrillConfig) IsAsync() bool {
	return mc.Async
}

func (mc MandrillConfig) ConnectionType() MailProviderConnectionType {
	return MailProviderConnectionTypeAPI
}

func (mc MandrillConfig) GetUsername() string {
	return ""
}

func (mc MandrillConfig) GetPassword() string {
	return mc.Key
}
