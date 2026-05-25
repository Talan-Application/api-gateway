package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Talan-Application/api-gateway/internal/ctxkeys"
)

const (
	ContextLocale    = "locale"
	defaultLocale    = "ru"
	headerXLocale    = "X-Locale"
	headerAcceptLang = "Accept-Language"
)

func LocaleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := resolveLocale(c)

		enriched := context.WithValue(c.Request.Context(), ctxkeys.LocaleKey, locale)
		c.Request = c.Request.WithContext(enriched)
		c.Set(ContextLocale, locale)

		c.Next()
	}
}

func resolveLocale(c *gin.Context) string {
	if v := strings.TrimSpace(c.GetHeader(headerXLocale)); v != "" {
		return strings.ToLower(v)
	}

	if v := c.GetHeader(headerAcceptLang); v != "" {
		// Accept-Language: kk  →  "kk"
		primary := strings.SplitN(v, ",", 2)[0]
		primary = strings.SplitN(primary, ";", 2)[0]
		primary = strings.SplitN(primary, "-", 2)[0]
		primary = strings.ToLower(strings.TrimSpace(primary))
		if primary != "" {
			return primary
		}
	}

	return defaultLocale
}
