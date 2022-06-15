package mailing

type MailProviderName string

func (mpn MailProviderName) String() string {
	return string(mpn)
}

const (
	MailProviderFile     MailProviderName = "file"
	MailProviderMailgun  MailProviderName = "mailgun"
	MailProviderMandrill MailProviderName = "mandrill"
	MailProviderSendgrid MailProviderName = "sendgrid"
	MailProviderSMTP     MailProviderName = "smtp"
)

type MailProviderConnectionType string

const (
	MailProviderConnectionTypeNone MailProviderConnectionType = "none"
	MailProviderConnectionTypeAPI  MailProviderConnectionType = "api"
	MailProviderConnectionTypeSMTP MailProviderConnectionType = "smtp"
)

type MailProviderEncryption string

const (
	MailProviderEncryptionNone     MailProviderEncryption = "none"
	MailProviderEncryptionSSL      MailProviderEncryption = "ssl"
	MailProviderEncryptionTLS      MailProviderEncryption = "tls"
	MailProviderEncryptionSSLTLS   MailProviderEncryption = "ssl/tls"
	MailProviderEncryptionSTARTTLS MailProviderEncryption = "starttls"
)
