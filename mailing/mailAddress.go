package mailing

import (
	"net/mail"
	"strings"

	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type MailAddress struct {
	Email string `json:"email" yaml:"email" valid:"required,email"`
	Name  string `json:"name" yaml:"name" valid:"required"`
}

func NewMailAddress(email, name string) MailAddress {
	return MailAddress{Email: email, Name: name}
}

func NewMailAddressFromInterface(i MailAddressInterface) MailAddress {
	return MailAddress{Email: i.GetEmail(), Name: i.GetName()}
}

func (ma MailAddress) IsEmpty() bool {
	return len(ma.Email) == 0
}

func (ma MailAddress) GetEmail() string {
	return ma.Email
}

func (ma MailAddress) GetName() string {
	return ma.Name
}

func (ma MailAddress) GetDomain() string {
	strs := strings.Split(ma.Email, "@")

	if len(strs) == 1 {
		return ""
	}

	return strs[1]
}

func (ma MailAddress) String() string {
	if len(ma.Email) == 0 {
		return ""
	}

	return (&mail.Address{
		Name:    ma.Name,
		Address: ma.Email,
	}).String()
}

func (ma MailAddress) Validate() (bool, error) {
	return contracts.ConfigValidate(ma)
}
