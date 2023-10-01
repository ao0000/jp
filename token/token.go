package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL" // ILLEGAL
	EOF               = "EOF"     // EOF

	STRING = "STRING" // STRING
	NUMBER = "NUMBER" // NUMBER

	NULL  = "NULL"  // null
	TRUE  = "TRUE"  // true
	FALSE = "FALSE" // false

	LBRACK = "LEFT_BRACKET"  // [
	LBRACE = "LEFT_BRACE"    // {
	RBRACK = "RIGHT_BRACKET" // ]
	RBRACE = "RIGHT_BRACE"   // }

	COLON = "COLON" // :
	COMMA = "COMMA" // ,

	ARRAY  = "ARRAY"  // ARRAY
	OBJECT = "OBJECT" // OBJECT
)

var ReservedWords = map[string]TokenType{
	"null":  NULL,
	"true":  TRUE,
	"false": FALSE,
}
