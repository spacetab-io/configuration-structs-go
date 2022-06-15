package mailing

import (
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type SendgridConfig struct {
	Key            string        `yaml:"key" valid:"required"`
	DkimPrivateKey string        `yaml:"dkimPrivateKey"`
	SendTimeout    time.Duration `yaml:"sendTimeout"`
}

func (sc SendgridConfig) String() string {
	return sc.Name().String()
}

func (sc SendgridConfig) Name() MailProviderName {
	return MailProviderSendgrid
}

func (sc SendgridConfig) IsAsync() bool {
	return true
}

func (sc SendgridConfig) GetHostPort() contracts.AddressInterface {
	return nil
}

func (sc SendgridConfig) GetEncryption() MailProviderEncryption {
	return MailProviderEncryptionNone
}

func (sc SendgridConfig) GetAuthType() contracts.AuthType {
	return contracts.AuthTypeNone
}

func (sc SendgridConfig) GetDKIMPrivateKey() *string {
	return &sc.DkimPrivateKey
}

func (sc SendgridConfig) GetConnectionTimeout() time.Duration {
	return 0
}

func (sc SendgridConfig) GetSendTimeout() time.Duration {
	return sc.SendTimeout
}

func (sc SendgridConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(sc)
}

func (sc SendgridConfig) ConnectionType() MailProviderConnectionType {
	return MailProviderConnectionTypeAPI
}

func (sc SendgridConfig) GetUsername() string {
	return ""
}

func (sc SendgridConfig) GetPassword() string {
	return sc.Key
}
