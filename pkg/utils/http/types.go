package http

const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	HeaderRequestID     = "X-Request-ID"
)

const (
	AuthSchemeBasic  = "Basic"
	AuthSchemeBearer = "Bearer"
)

const (
	MIMEApplicationJSON = "application/json"
	MIMETextEventStream = "text/event-stream"
)

const (
	TokenCookieName        = "token"
	RefreshTokenCookieName = "akuity.refresh-token"
	ArgoCDTokenCookieName  = "argocd.token"
)
