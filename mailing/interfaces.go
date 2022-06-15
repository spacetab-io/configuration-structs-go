package mailing

import (
	"fmt"
	"time"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MailsConfigInterface interface {
	contracts.ValidatableInterface
	fmt.Stringer

	GetActiveProviderConfig() (MailProviderConfigInterface, error)
	GetMessageConfig() MessagingConfigInterface
}

type MessagingConfigInterface interface {
	contracts.ValidatableInterface
	fmt.Stringer

	GetFrom() MailAddressInterface
	GetReplyTo() MailAddressInterface
	GetSubjectPrefix() string
}

type MailAddressInterface interface {
	contracts.ValidatableInterface
	fmt.Stringer

	IsEmpty() bool
	GetEmail() string
	GetName() string
	GetDomain() string
}

type MailAddressListInterface interface {
	contracts.ValidatableInterface

	IsEmpty() bool
	GetStringList() []string
	GetList() []MailAddressInterface
}

type MailProviderConfigInterface interface {
	contracts.ValidatableInterface
	fmt.Stringer

	Name() MailProviderName
	IsAsync() bool
	ConnectionType() MailProviderConnectionType
	GetUsername() string
	GetPassword() string
	GetHostPort() contracts.AddressInterface
	GetEncryption() MailProviderEncryption
	GetAuthType() contracts.AuthType
	GetDKIMPrivateKey() *string
	GetConnectionTimeout() time.Duration
	GetSendTimeout() time.Duration
}
