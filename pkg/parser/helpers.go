package parser

import (
	"errors"
	"fmt"
	"github.com/reilandeubank/golisp/pkg/scanner"
)

// match parses a token and if it matches any of the passed TokenTypes
// it advances the parser (critical) and returns true
func (p *Parser) match(types ...scanner.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

// check takes a peek at the current token and returns true if its type matches passed type t
func (p *Parser) check(t scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

// advance increments the parser and returns the previous token
func (p *Parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.Curr++
	}
	return p.previous()
}

// isAtEnd checks for EOF
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == scanner.EOF
}

// peek returns the current token without advancing the parser
func (p *Parser) peek() scanner.Token {
	return p.Tokens[p.Curr]
}

// previous returns the token that was just passed
func (p *Parser) previous() scanner.Token {
	return p.Tokens[p.Curr-1]
}

// consume will advance past a specified TokenType or return an error otherwise
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

// stringify returns string representation of passed object
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
