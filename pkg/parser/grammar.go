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

// func (p *Parser) equality() (Expression, error) {
// 	expr, err := p.comparison()
// 	for p.match(scanner.EQUAL) {
// 		operator := p.previous()
// 		right, err := p.comparison()
// 		if err != nil {
// 			return Atom{Value: nil}, err
// 		}
// 		expr = Binary{Left: expr, Operator: operator, Right: right}
// 	}
// 	return expr, err
// }

// func (p *Parser) comparison() (Expression, error) {
// 	expr, err := p.term()
// 	for p.match(scanner.GREATER, scanner.LESS) {
// 		operator := p.previous()
// 		right, err := p.term()
// 		if err != nil {
// 			return Atom{Value: nil}, err
// 		}
// 		expr = Binary{Left: expr, Operator: operator, Right: right}
// 	}
// 	return expr, err
// }

// func (p *Parser) term() (Expression, error) {
//     // First, check if the current token is an operator
//     if !p.match(scanner.PLUS, scanner.MINUS) {
//         return Atom{Value: nil}, errors.New("expected operator (+ or -)")
//     }

//     // Get the operator
//     operator := p.previous()

//     // Parse the first operand
//     left, err := p.factor()
//     if err != nil {
//         return Atom{Value: nil}, err
//     }

//     // Parse the second operand
//     right, err := p.factor()
//     if err != nil {
//         return Atom{Value: nil}, err
//     }

//     // Return the binary expression
//     return Binary{Left: left, Operator: operator, Right: right}, nil
// }

// func (p *Parser) factor() (Expression, error) {
// 	 // First, check if the current token is an operator
// 	 if !p.match(scanner.SLASH, scanner.STAR) {
//         return Atom{Value: nil}, errors.New("expected operator (* or /)")
//     }

//     // Get the operator
//     operator := p.previous()

//     // Parse the first operand
//     left, err := p.unary()
//     if err != nil {
//         return Atom{Value: nil}, err
//     }

//     // Parse the second operand
//     right, err := p.unary()
//     if err != nil {
//         return Atom{Value: nil}, err
//     }

//     // Return the binary expression
//     return Binary{Left: left, Operator: operator, Right: right}, nil
// }

// func (p *Parser) unary() (Expression, error) {
// 	if p.match(scanner.MINUS) {
// 		operator := p.previous()
// 		right, err := p.unary()
// 		if err != nil {
// 			return Atom{Value: nil}, err
// 		}
// 		return Unary{Operator: operator, Right: right}, err
// 	}
// 	return p.primary()
// }

// func (p *Parser) primary() (Expression, error) {
// 	if p.match(scanner.LEFT_PAREN) {
// 		if !p.match(scanner.PLUS, scanner.MINUS, scanner.STAR, scanner.SLASH, /* other operators */) {
//             return Atom{Value: nil}, errors.New("expected operator")
//         }
//         operator := Operator{Operator: p.previous()}

// 		var operands []Expression
// 		for !p.check(scanner.RIGHT_PAREN) && !p.isAtEnd() {
// 			operand, err := p.expr() // recursive call to handle nested expressions
// 			if err != nil {
// 				return Atom{Value: nil}, err
// 			}
// 			operands = append(operands, operand)
// 		}

// 		_, err := p.consume(scanner.RIGHT_PAREN, "expect ')' after expression")
// 		if err != nil {
// 			return Atom{Value: nil}, err
// 		}

// 		return ListExpr{Head: operator, Tail: operands}, nil
// 	}

// 	if p.match(scanner.NUMBER) {
//         // Handle numbers
//         return Atom{Value: p.previous().Literal}, nil
//     }

// 	if p.match(scanner.STRING) {
//         // Handle strings
//         return Atom{Value: p.previous().Literal}, nil
//     }

// 	if p.match(scanner.SYMBOL) {
//         // Handle symbols (identifiers)
//         return Atom{Value: p.previous().Literal}, nil
//     }

// 	if p.match(scanner.FALSE) {
//         // Handle false boolean literal
//         return Atom{Value: false}, nil
//     }

// 	if p.match(scanner.TRUE) {
//         // Handle true boolean literal
//         return Atom{Value: true}, nil
//     }

//     if p.match(scanner.NIL) {
//         // Handle nil
//         return Atom{Value: nil}, nil
//     }

// 	return Atom{Value: nil}, errors.New("unexpected token")
// }

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
		// Handle symbols (identifiers)
		return Atom{Value: p.previous().Literal}, nil
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