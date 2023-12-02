package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
)

func ParseError(t scanner.Token, message string) {
	if t.Type == scanner.EOF {
		scanner.Report(t.Line, " at end", message)
	} else {
		scanner.Report(t.Line, " at '" + t.Lexeme + "'", message)
	}
}

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
