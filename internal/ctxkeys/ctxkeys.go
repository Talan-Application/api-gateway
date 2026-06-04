package ctxkeys

type contextKey string

const (
	LocaleKey     contextKey = "locale"
	AuthHeaderKey contextKey = "auth_header"
)
