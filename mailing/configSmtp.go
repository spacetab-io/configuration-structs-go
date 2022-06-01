package mailing

import (
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type SMTPConfig struct {
	Enabled           bool                   `yaml:"enabled" valid:"optional"`
	Host              string                 `yaml:"host" valid:"required"`
	Port              uint                   `yaml:"port" valid:"required"`
	Username          string                 `yaml:"username" valid:"optional"`
	Password          string                 `yaml:"password" valid:"optional"`
	Encryption        MailProviderEncryption `yaml:"encryption" valid:"required"`
	AuthType          contracts.AuthType     `yaml:"authType" valid:"required"`
	DKIMPrivateKey    string                 `yaml:"DKIMPrivateKey" valid:"optional"`
	ConnectionTimeout time.Duration          `yaml:"connectionTimeout" valid:"required"`
	SendTimeout       time.Duration          `yaml:"sendTimeout" valid:"required"`
}

func (sc SMTPConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(sc)
}

func (sc SMTPConfig) String() string {
	return "smtp"
}

func (sc SMTPConfig) IsEnable() bool {
	return sc.Enabled
}

func (sc SMTPConfig) IsAsync() bool {
	return false
}

func (sc SMTPConfig) ConnectionType() MailProviderConnectionType {
	return MailProviderConnectionTypeSMTP
}

func (sc SMTPConfig) GetUsername() string {
	return sc.Username
}

func (sc SMTPConfig) GetPassword() string {
	return sc.Password
}

func (sc SMTPConfig) GetHostPort() contracts.AddressInterface {
	return &contracts.HostCfg{Host: sc.Host, Port: sc.Port}
}

func (sc SMTPConfig) GetEncryption() MailProviderEncryption {
	return sc.Encryption
}

func (sc SMTPConfig) GetAuthType() contracts.AuthType {
	return sc.AuthType
}

func (sc SMTPConfig) GetDKIMPrivateKey() *string {
	return &sc.DKIMPrivateKey
}

func (sc SMTPConfig) GetConnectionTimeout() time.Duration {
	return sc.ConnectionTimeout
}

func (sc SMTPConfig) GetSendTimeout() time.Duration {
	return sc.SendTimeout
}
