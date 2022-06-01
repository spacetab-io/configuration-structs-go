package mailing

type MailProviderName string

func (mpn MailProviderName) String() string {
	return string(mpn)
}

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
