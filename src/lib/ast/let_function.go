package ast

import "github.com/raviqqe/tisp/src/lib/debug"

type LetFunction struct {
	name      string
	signature Signature
	lets      []interface{}
	body      interface{}
	info      debug.Info
}

func NewLetFunction(name string, sig Signature, lets []interface{}, expr interface{}, i debug.Info) LetFunction {
	return LetFunction{name, sig, lets, expr, i}
}

func (f LetFunction) Name() string {
	return f.name
}

func (f LetFunction) Signature() Signature {
	return f.signature
}

// Lets returns let statements contained in this let-function statement.
// Return values should be LetVar or LetFunction.
func (f LetFunction) Lets() []interface{} {
	return f.lets
}

func (f LetFunction) Body() interface{} {
	return f.body
}

func (f LetFunction) DebugInfo() debug.Info {
	return f.info
}