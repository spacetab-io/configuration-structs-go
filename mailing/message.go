package mailing

import (
	"strings"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
	"github.com/spacetab-io/configuration-structs-go/v2/mime"
)

type MessagingConfig struct {
	From          MailAddress `yaml:"from" valid:"required"`
	ReplyTo       MailAddress `yaml:"replyTo" valid:"required"`
	MimeType      mime.Type   `yaml:"mimeType" valid:"required,in(text/plain|text/html)"`
	SubjectPrefix string      `yaml:"subjectPrefix" valid:"optional"`
}

func (mc MessagingConfig) GetSubjectPrefix() string {
	return mc.SubjectPrefix
}

func (mc MessagingConfig) String() string {
	addresses := make([]string, 0, 3) // nolint: gomnd

	for _, v := range []struct {
		name    string
		address MailAddress
	}{
		{name: "from", address: mc.From},
		{name: "reply-to", address: mc.ReplyTo},
	} {
		if !v.address.IsEmpty() {
			addresses = append(addresses, v.name+": "+v.address.String())
		}
	}

	return strings.Join(addresses, ", ")
}

func (mc MessagingConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(mc)
}

func (mc MessagingConfig) GetFrom() MailAddressInterface {
	return mc.From
}

func (mc MessagingConfig) GetReplyTo() MailAddressInterface {
	return mc.ReplyTo
}
