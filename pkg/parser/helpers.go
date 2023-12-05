package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
	"errors"
	"fmt"
)

func (p *Parser) match(types ...scanner.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.Curr++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == scanner.EOF
}

func (p *Parser) peek() scanner.Token {
	return p.Tokens[p.Curr]
}

func (p *Parser) previous() scanner.Token {
	return p.Tokens[p.Curr-1]
}

func (p *Parser) consume(t scanner.TokenType, message string) (scanner.Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}
	ParseError(p.peek(), message)
	return scanner.NewToken(scanner.OTHER, "", nil, 0), errors.New(message)
}

func (p *Parser) isKeyword() bool {
	return p.match(scanner.DEFINE, scanner.SET, scanner.CONS, scanner.COND, scanner.CAR, scanner.CDR, scanner.NIL, scanner.TRUE, scanner.FALSE, scanner.ANDQ, scanner.ORQ, scanner.NOTQ, scanner.NUMBERQ, scanner.SYMBOLQ, scanner.LISTQ, scanner.NILQ)
}

func stringify(object interface{}) string {
	if object == nil {
		return "nil"
	}

	// Type assertion for float64
	if val, ok := object.(float64); ok {
		return fmt.Sprintf("%g", val) // %g removes trailing zeros
	}

	// Default to using fmt.Sprintf for other types
	return fmt.Sprintf("%v", object)
}