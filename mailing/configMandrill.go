package mailing

import (
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MandrillConfig struct {
	Enabled bool   `yaml:"enabled" valid:"optional"`
	Async   bool   `yaml:"async" valid:"optional,bool"`
	Key     string `yaml:"key" valid:"required"`
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

func (mc MandrillConfig) String() string {
	return "mandrillAPI"
}

func (mc MandrillConfig) IsEnable() bool {
	return mc.Enabled
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
