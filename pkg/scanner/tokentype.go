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
	EQUAL
	GREATER
	LESS

	// Literals.
	SYMBOL
	STRING
	NUMBER

	// Keywords.
	DEFINE
	SET
	CONS
	COND
	CAR
	CDR
	NIL
	TRUE
	FALSE
	ANDQ
	ORQ
	NOTQ
	NUMBERQ
	SYMBOLQ
	LISTQ
	NILQ

	WHITESPACE
	OTHER
	EOF
)
