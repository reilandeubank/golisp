package parser

import (
	"errors"
	"golisp/pkg/scanner"
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

		// If Head is a symbol, evaluate and return function call
		if funcName, ok := head.(Symbol); ok {
			return p.functionCall(funcName)
		}

		// If Head is 'define', we expect a function definition and return it
		if kw, ok := head.(Keyword); ok && kw.Keyword.Type == scanner.DEFINE {
			return p.functionDefinition()
		}

		// If the list isn't a function call or definition, it is a normal list
		// so we will simply evaluate each element and build up the tail
		var tail []Expression
		for !p.check(scanner.RIGHT_PAREN) && !p.isAtEnd() {
			expr, err := p.expr() // Parse each operand
			if err != nil {
				return nil, err
			}
			tail = append(tail, expr)
		}

		// Ensure the list ends with right parentheses
		_, err = p.consume(scanner.RIGHT_PAREN, "expect ')' after expression")
		if err != nil {
			return nil, err
		}

		return ListExpr{Head: head, Tail: tail}, nil
	}

	// If it's not a list, it might be an atom or other type of expression
	return p.atom()
}

func (p *Parser) functionCall(funcName Symbol) (Expression, error) {
	// Expecting a list of parameters
	params, err := p.callParamList()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(scanner.RIGHT_PAREN, "Expect ')' after function call")
	if err != nil {
		return nil, err
	}

	return Call{Callee: funcName, Token: funcName.Name, ArgsList: params}, nil
}

func (p *Parser) callParamList() ([]Expression, error) {
	var params []Expression

	// build a list of evaluated parameters
	for !p.check(scanner.RIGHT_PAREN) && !p.isAtEnd() {
		param, err := p.expr()
		if err != nil {
			return nil, err
		}
		params = append(params, param)
	}

	return params, nil

}

func (p *Parser) functionDefinition() (Expression, error) {
	functionName, err := p.consume(scanner.SYMBOL, "Expect function name.")
	if err != nil {
		return nil, err
	}

	// Expecting a list of parameters
	params, err := p.paramList()
	if err != nil {
		return nil, err
	}

	// The function body is an expression
	body, err := p.expr()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(scanner.RIGHT_PAREN, "Expect ')' after function definition.")
	if err != nil {
		return nil, err
	}

	// Function definition is returned but will not be printed to terminal like other expressions
	return FuncDefinition{Name: functionName, Params: params, Body: body}, nil
}

func (p *Parser) paramList() ([]scanner.Token, error) {
	var params []scanner.Token

	_, err := p.consume(scanner.LEFT_PAREN, "Expect '(' after function name.")
	if err != nil {
		return nil, err
	}

	// Build a list of parameter names that will be assigned to the passed parameters at call-time
	for !p.match(scanner.RIGHT_PAREN) && !p.isAtEnd() {
		param, err := p.consume(scanner.SYMBOL, "Expect parameter name.")
		if err != nil {
			return nil, err
		}
		params = append(params, param)
	}

	return params, nil

}

func (p *Parser) atom() (Expression, error) {
	// handles keywords as the first element of a list
	// Native functions are probably a better way to do this
	if p.isKeyword() {
		k := Keyword{Keyword: p.previous()}
		return k, nil
	}

	// Works similarly to keywords. Native functions might've been better
	if p.match(scanner.PLUS, scanner.MINUS, scanner.STAR, scanner.SLASH, scanner.EQUAL, scanner.LESS, scanner.GREATER) {
		// Handle operators
		return Operator{Operator: p.previous()}, nil
	}

	// Simply returning an atom when encountered. Only handles numbers and strings
	if p.match(scanner.NUMBER, scanner.STRING) {
		var prevValue interface{} = p.previous().Literal
		var err error
		switch prevValue.(type) {
		case string:
			return Atom{Value: prevValue, Type: scanner.STRING}, err
		case float64:
			return Atom{Value: prevValue, Type: scanner.NUMBER}, err
		default:
			// Handle other types or error
			message := "unexpected literal type: " + fmt.Sprintf("%T", prevValue)
			ParseError(p.peek(), message)
			err = errors.New(message)
		}
		return Atom{Value: nil, Type: scanner.NIL}, err
	}

	// Symbols, TRUE, and NIL are returned mostly as-is
	if p.match(scanner.SYMBOL) {
		return Symbol{Name: p.previous()}, nil
	}

	if p.match(scanner.TRUE) {
		// Handle true boolean literal
		return Atom{Value: true}, nil
	}

	if p.match(scanner.NIL) {
		// Handle nil
		return Atom{Value: nil}, nil
	}

	// If none of the above get triggered, something went wrong
	return Atom{Value: nil}, errors.New("unexpected token: " + p.peek().Lexeme)
}
