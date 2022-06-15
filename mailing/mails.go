package mailing

import (
	"errors"
	"fmt"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type AvailableProvidersConfig struct {
	File     FileConfig     `yaml:"file"`
	Mailgun  MailgunConfig  `yaml:"mailgun"`
	Mandrill MandrillConfig `yaml:"mandrill"`
	Sendgrid SendgridConfig `yaml:"sendgrid"`
	SMTP     SMTPConfig     `yaml:"smtp"`
}

type ProvidersConfig struct {
	Active    MailProviderName         `yaml:"active" valid:"required,in(file|mailgun|mandrill|sendgrid|smtp)"`
	Available AvailableProvidersConfig `yaml:"available"`
}

type MailsConfig struct {
	Message   MessagingConfig `yaml:"message"`
	Providers ProvidersConfig `yaml:"providers"`
}

func (mmc MailsConfig) String() string {
	cfg, err := mmc.GetActiveProviderConfig()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%s: %s", mmc.Providers.Active.String(), cfg.String())
}

func (mmc MailsConfig) Validate() (bool, error) {
	cfg, err := mmc.GetActiveProviderConfig()
	if err != nil {
		return false, err
	}

	ok, err := contracts.ConfigValidate(mmc.Message)
	if !ok || err != nil {
		return ok, err
	}

	return contracts.ConfigValidate(cfg)
}

var ErrUnknownProvider = errors.New("unknown email provider")

func (mmc MailsConfig) GetActiveProviderConfig() (MailProviderConfigInterface, error) {
	var cfg MailProviderConfigInterface

	switch mmc.Providers.Active {
	case MailProviderFile:
		cfg = &mmc.Providers.Available.File
	case MailProviderMailgun:
		cfg = &mmc.Providers.Available.Mailgun
	case MailProviderMandrill:
		cfg = &mmc.Providers.Available.Mandrill
	case MailProviderSendgrid:
		cfg = &mmc.Providers.Available.Sendgrid
	case MailProviderSMTP:
		cfg = &mmc.Providers.Available.SMTP
	default:
		return nil, ErrUnknownProvider
	}

	return cfg, nil
}

func (mmc MailsConfig) GetMessageConfig() MessagingConfigInterface {
	return mmc.Message
}
