package tracing

import (
	"github.com/lucsky/cuid"
)

type Context struct {
	RequestID     string
	RequestSource string
}

// NewContext creates a new tracing context
func NewContext(gitHash string) *Context {
	return &Context{
		RequestID:     cuid.New(),
		RequestSource: "lernen-api",
	}
}

//GetOutgoingHeaders
func (t *Context) GetOutgoingHeaders() map[string]string {
	return map[string]string{
		"X-Request-ID":     t.RequestID,
		"X-Request-Source": t.RequestSource,
	}
}
