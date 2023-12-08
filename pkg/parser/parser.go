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

func (p *Parser) Parse() ([]Expression, error) {
	var expressions []Expression

	for !p.isAtEnd() {
		expr, err := p.expr()
		if err != nil {
			return []Expression{}, err
		}
		expressions = append(expressions, expr)
	}

	return expressions, nil
}