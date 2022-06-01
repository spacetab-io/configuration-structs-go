package contracts

type AuthType string

// http auth types.
const (
	AuthTypeBasic AuthType = "basic"
	AuthTypeJWT   AuthType = "jwt"
)

// email server auth types.
const (
	AuthTypeNone    AuthType = "none"
	AuthTypePlain   AuthType = "plain"
	AuthTypeLogin   AuthType = "login"
	AuthTypeCRAMMD5 AuthType = "cram-md5"
)
