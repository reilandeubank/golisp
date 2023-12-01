package scanner

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	SYMBOL
	STRING
	NUMBER
	SEXPR

	// Keywords.
	DEFINE
	SET
	CONS
	COND
	CAR
	CDR
	ANDQ
	ORQ
	NOTQ
	NUMBERQ
	SYMBOLQ
	LISTQ
	NILQ
	EQQ

	WHITESPACE
	OTHER
	EOF
)
