package ir

import "github.com/tisp-lang/tisp/src/lib/core"

// Case represents a case of a pattern and a corresponding value in a switch expression.
type Case struct {
	pattern *core.Thunk
	value   int
}

// NewCase creates a case in a switch expression.
func NewCase(p *core.Thunk, v int) Case {
	return Case{p, v}
}
