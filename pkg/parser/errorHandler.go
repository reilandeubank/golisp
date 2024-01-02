package parser

import (
	"golisp/pkg/scanner"
)

// ParseError reports an error through the scanner. Probably should've used the same
// error handling from the interpreter, but this was already built
func ParseError(t scanner.Token, message string) {
	if t.Type == scanner.EOF {
		scanner.Report(t.Line, " at end", message)
	} else {
		scanner.Report(t.Line, " at '"+t.Lexeme+"'", message)
	}
}

// synchronize will advance the parser until it reaches the end of the current expression
// to get back to a known state after an error is encountered
func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == scanner.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case scanner.DEFINE:
		case scanner.SET:
		case scanner.COND:
		}

		p.advance()
	}
}
