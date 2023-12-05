package parser

import (
	"github.com/reilandeubank/golisp/pkg/scanner"
	"errors"
	// "strconv"
	"fmt"
)

func (p *Parser) expr() (Expression, error) {
	return p.list()
}

func (p *Parser) list() (Expression, error) {
    if p.match(scanner.LEFT_PAREN) {
        // handle parsing of list
        head, err := p.expr() // First element is the operator or function
        if err != nil {
            return nil, err
        }

        var tail []Expression
        for !p.check(scanner.RIGHT_PAREN) && !p.isAtEnd() {
            if p.check(scanner.LEFT_PAREN) {
                // Handle nested list
                nestedList, err := p.list() // Evaluate the nested list
                if err != nil {
                    return nil, err
                }
                tail = append(tail, nestedList) // Append the evaluated nested list as a single expression
            } else {
                expr, err := p.expr() // Parse each operand
                if err != nil {
                    return nil, err
                }
                tail = append(tail, expr)
            }
        }

        _, err = p.consume(scanner.RIGHT_PAREN, "expect ')' after expression")
        if err != nil {
            return nil, err
        }

        return ListExpr{Head: head, Tail: tail}, nil
    }

    // If it's not a list, it might be an atom or other type of expression
    return p.atom()
}

func (p *Parser) atom() (Expression, error) {
	if p.isKeyword() {
		k := Keyword{Keyword: p.previous()}
		// fmt.Println("adding keyword:", scanner.KeywordsReverse[k.Keyword.Type])
		return k, nil
	}

	if p.match(scanner.PLUS, scanner.MINUS, scanner.STAR, scanner.SLASH, scanner.EQUAL, scanner.LESS, scanner.GREATER) {
		// Handle operators
		return Operator{Operator: p.previous()}, nil
	}

	if p.match(scanner.NUMBER, scanner.STRING) {
		var prevValue interface{} = p.previous().Literal
		var err error
		switch prevValue.(type) {
		case string:
			// fmt.Println("String: " + prevValue.(string))
			return Atom{Value: prevValue, Type: scanner.STRING}, err
		case float64:
			// fmt.Println("Number: " + fmt.Sprintf("%f", prevValue.(float64)))
			return Atom{Value: prevValue, Type: scanner.NUMBER}, err
		default:
			// Handle other types or error
			message := "unexpected literal type: " + fmt.Sprintf("%T", prevValue)
			ParseError(p.peek(), message)
			err = errors.New(message)
		}
		return Atom{Value: nil, Type: scanner.NIL}, err
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