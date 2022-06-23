package mime

type Type string

const (
	TextHTML  Type = "text/html"
	TextPlain Type = "text/plain"
)

func (mt Type) String() string {
	return string(mt)
}

func (mt Type) IsEmpty() bool {
	return mt == ""
}
