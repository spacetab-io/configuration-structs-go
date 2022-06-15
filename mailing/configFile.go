package mailing

import (
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type FileConfig struct {
	FilePath string `yaml:"filePath" valid:"required"`
	Async    bool   `yaml:"isAsync" valid:"-"`
}

func (fc FileConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(fc)
}

func (fc FileConfig) String() string {
	return fc.Name().String()
}

func (fc FileConfig) Name() MailProviderName {
	return MailProviderFile
}

func (fc FileConfig) IsAsync() bool {
	return fc.Async
}

func (fc FileConfig) ConnectionType() MailProviderConnectionType {
	return MailProviderConnectionTypeNone
}

func (fc FileConfig) GetUsername() string {
	return ""
}

func (fc FileConfig) GetPassword() string {
	return ""
}

func (fc FileConfig) GetHostPort() contracts.AddressInterface {
	return &contracts.HostCfg{Host: fc.FilePath}
}

func (fc FileConfig) GetEncryption() MailProviderEncryption {
	return MailProviderEncryptionNone
}

func (fc FileConfig) GetAuthType() contracts.AuthType {
	return contracts.AuthTypeNone
}

func (fc FileConfig) GetDKIMPrivateKey() *string {
	return nil
}

func (fc FileConfig) GetConnectionTimeout() time.Duration {
	return 0
}

func (fc FileConfig) GetSendTimeout() time.Duration {
	return 0
}
