package interpreter

import (
	"fmt"

	"github.com/reilandeubank/golisp/pkg/scanner"
)

// RuntimeError defines a new Error type. Did not know this was possible until I started work on
// the interpreter, so this would be the better way to implement my scanner and parser errors
type RuntimeError struct {
	Token   scanner.Token
	Message string
}

func (r *RuntimeError) Error() string {
	msg := fmt.Sprintf("[line %d] Runtime Error: %s", r.Token.Line, r.Message)
	return msg
}
