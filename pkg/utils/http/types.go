package http

const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	HeaderPlatform      = "X-Platform"
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
	TokenCookieName            = "token"
	RefreshTokenCookieName     = "akuity.refresh-token"
	PlatformCookieName         = "x-platform"
	ArgoCDTokenCookieName      = "argocd.token"
	KargoTokenCookieName       = "kargo.token"
	AimsTokenCookieName        = "aims.token"
	AimsRefreshTokenCookieName = "aims.refresh-token"
)
