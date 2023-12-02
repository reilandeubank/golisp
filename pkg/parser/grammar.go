package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
	"errors"
	"strconv"
	"fmt"
)

func (p *Parser) expr() (Expression, error) {
	return p.list()
}

func (p *Parser) list() (Expression, error) {
    if p.match(scanner.LEFT_PAREN) {
        // handle parsing of list
        operator, err := p.expr() // First element is the operator or function
        if err != nil {
            return nil, err
        }

        var operands []Expression
        for !p.check(scanner.RIGHT_PAREN) && !p.isAtEnd() {
            operand, err := p.expr() // Parse each operand
            if err != nil {
                return nil, err
            }
            operands = append(operands, operand)
        }

        _, err = p.consume(scanner.RIGHT_PAREN, "expect ')' after expression")
        if err != nil {
            return nil, err
        }

        return ListExpr{Head: operator, Tail: operands}, nil
    }

    // If it's not a list, it might be an atom or other type of expression
    return p.atom()
}

func (p *Parser) atom() (Expression, error) {
	if p.isKeyword() {
		return Keyword{Keyword: p.peek()}, nil
	}

	if p.match(scanner.PLUS, scanner.MINUS, scanner.STAR, scanner.SLASH, scanner.EQUAL, scanner.LESS, scanner.GREATER) {
		// Handle operators
		return Operator{Operator: p.previous()}, nil
	}

	if p.match(scanner.NUMBER, scanner.STRING) {
		var value string
		var err error
		switch v := p.previous().Literal.(type) {
		case string:
			value = v
		case float64:
			value = strconv.FormatFloat(v, 'f', -1, 64)
		default:
			message := fmt.Sprintf("unexpected literal type %T", v)
			ParseError(p.peek(), message)
			err = errors.New(message)
		}
		return Atom{Value: value}, err
	}

	if p.match(scanner.SYMBOL) {
		// fmt.Println("Symbol: " + p.previous().Lexeme)
		return Symbol{Name: p.previous()}, nil
	} 

	if p.match(scanner.FALSE) {
		// Handle false boolean literal
		return Atom{Value: false}, nil
	}

	if p.match(scanner.TRUE) {
		// Handle true boolean literal
		return Atom{Value: true}, nil
	}

	if p.match(scanner.NIL) {
		// Handle nil
		return Atom{Value: nil}, nil
	}

	return Atom{Value: nil}, errors.New("unexpected token: " + p.peek().Lexeme)
}