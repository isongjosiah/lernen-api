package common

import "fmt"

type ContentType int

const (
	ContentTypeJSON ContentType = 0
)

// ContextKey ...
type ContextKey string

func (c ContextKey) String() string {
	return fmt.Sprintf("lernen-api-context-key-%v", string(c))
}
