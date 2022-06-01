package mailing

import (
	"strings"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MailingsConfig struct {
	From          MailAddress `yaml:"from" valid:"required"`
	ReplyTo       MailAddress `yaml:"replyTo" valid:"required"`
	SubjectPrefix string      `yaml:"subjectPrefix" valid:"optional"`
}

func (mc MailingsConfig) GetSubjectPrefix() string {
	return mc.SubjectPrefix
}

func (mc MailingsConfig) String() string {
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

func (mc MailingsConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(mc)
}

func (mc MailingsConfig) GetFrom() MailAddressInterface {
	return mc.From
}

func (mc MailingsConfig) GetReplyTo() MailAddressInterface {
	return mc.ReplyTo
}
