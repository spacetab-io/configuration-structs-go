package cfgstructs

type ServiceName string

func (n ServiceName) String() string {
	return string(n)
}
