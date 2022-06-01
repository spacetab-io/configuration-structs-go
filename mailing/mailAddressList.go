package mailing

type MailAddressList []MailAddress

func NewMailAddressListFromInterface(il MailAddressListInterface) MailAddressList {
	mal := make([]MailAddress, 0, len(il.GetList()))

	for _, i := range il.GetList() {
		mal = append(mal, MailAddress{Email: i.GetEmail(), Name: i.GetName()})
	}

	return mal
}

func (mal MailAddressList) Validate() (bool, error) {
	for _, ma := range mal {
		ok, err := ma.Validate()
		if !ok {
			return ok, err
		}
	}

	return true, nil
}

func (mal MailAddressList) IsEmpty() bool {
	return len(mal) == 0
}

func (mal MailAddressList) GetStringList() []string {
	if len(mal) == 0 {
		return nil
	}

	strList := make([]string, 0, len(mal))

	for _, ma := range mal {
		strList = append(strList, ma.String())
	}

	return strList
}

func (mal MailAddressList) GetList() []MailAddressInterface {
	maList := make([]MailAddressInterface, 0, len(mal))

	for _, ma := range mal {
		maList = append(maList, ma)
	}

	return maList
}
