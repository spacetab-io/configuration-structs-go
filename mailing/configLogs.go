package mailing

import (
	"time"

	cfgstructs "github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type LogsConfig struct {
	Stream string `yaml:"stream" valid:"required,in(stdout|stderr)"`
}

func (m LogsConfig) Validate() (bool, error) {
	return true, nil
}

func (m LogsConfig) String() string {
	return m.Name().String()
}

func (m LogsConfig) Name() MailProviderName {
	return "mock"
}

func (m LogsConfig) IsEnable() bool {
	return true
}

func (m LogsConfig) IsAsync() bool {
	return false
}

func (m LogsConfig) ConnectionType() MailProviderConnectionType {
	return MailProviderConnectionTypeNone
}

func (m LogsConfig) GetUsername() string {
	return ""
}

func (m LogsConfig) GetPassword() string {
	return ""
}

func (m LogsConfig) GetHostPort() cfgstructs.AddressInterface {
	return &cfgstructs.HostCfg{Host: m.Stream}
}

func (m LogsConfig) GetEncryption() MailProviderEncryption {
	return MailProviderEncryptionNone
}

func (m LogsConfig) GetAuthType() cfgstructs.AuthType {
	return cfgstructs.AuthTypeNone
}

func (m LogsConfig) GetDKIMPrivateKey() *string {
	return nil
}

func (m LogsConfig) GetConnectionTimeout() time.Duration {
	return 0
}

func (m LogsConfig) GetSendTimeout() time.Duration {
	return 0
}
