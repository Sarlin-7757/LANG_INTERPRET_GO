package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" /* Signifies the token/character we don't know */
	EOF     = "EOF"     // end of file

	// Identifiers + literals

	IDENT = "IDENT" // add , foobar , x, y , ....
	INT   = "INT"   // 13245

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimitors
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
