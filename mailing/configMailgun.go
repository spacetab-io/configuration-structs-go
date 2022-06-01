package mailing

import (
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MailgunConfig struct {
	Enabled        bool          `yaml:"enabled"`
	APIBase        string        `yaml:"apiBase"`
	Domain         string        `yaml:"domain"`
	Key            string        `yaml:"key"`
	DKIMPrivateKey string        `yaml:"DKIMPrivateKey"`
	SendTimeout    time.Duration `yaml:"sendTimeout"`
}

func (mgc MailgunConfig) IsAsync() bool {
	return true
}

func (mgc MailgunConfig) GetHostPort() contracts.AddressInterface {
	var host string

	switch mgc.APIBase {
	case "us":
		host = mailgun.APIBaseUS
	case "eu":
		host = mailgun.APIBaseEU
	default:
		host = mailgun.APIBase
	}

	return &contracts.HostCfg{Host: host, Port: 443} // nolint: gomnd
}

func (mgc MailgunConfig) GetEncryption() MailProviderEncryption {
	return MailProviderEncryptionNone
}

func (mgc MailgunConfig) GetAuthType() contracts.AuthType {
	return contracts.AuthTypeNone
}

func (mgc MailgunConfig) GetDKIMPrivateKey() *string {
	return &mgc.DKIMPrivateKey
}

func (mgc MailgunConfig) GetConnectionTimeout() time.Duration {
	return time.Duration(0)
}

func (mgc MailgunConfig) GetSendTimeout() time.Duration {
	return time.Duration(0)
}

func (mgc MailgunConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(mgc)
}

func (mgc MailgunConfig) IsEnable() bool {
	return mgc.Enabled
}

func (mgc MailgunConfig) ConnectionType() MailProviderConnectionType {
	return MailProviderConnectionTypeAPI
}

func (mgc MailgunConfig) GetUsername() string {
	return mgc.Domain
}

func (mgc MailgunConfig) GetPassword() string {
	return mgc.Key
}

func (mgc MailgunConfig) String() string {
	return "mailgunAPI"
}
