package contracts

const ContextKeyRequestID ContextKey = "RequestID"

type ContextKey string

func (k ContextKey) String() string {
	return string(k)
}
