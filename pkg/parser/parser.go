package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
)

type Parser struct {
	Tokens []scanner.Token
	Curr int
}

func NewParser(tokens []scanner.Token) Parser {
	return Parser{
		Tokens: tokens,
		Curr: 0,
	}
}

func (p *Parser) Parse() (Expression, error) {
	// var statements []Stmt

	// for !p.isAtEnd() {
	// 	dec, err := p.declaration()
	// 	if err != nil {
	// 		return []Stmt{}, err
	// 	}
	// 	statements = append(statements, dec)
	// }

	// return statements, nil

	expr, err := p.expr()
	if err != nil {
		return Atom{Value: nil}, err
	}
	return expr, err
}