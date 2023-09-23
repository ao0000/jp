package lexer

import (
	"github.com/ao0000/jp/token"
)

type lexer struct {
	rawJSON   string
	index     int
	nextIndex int
}

// TODO: Extract lexer's behavior to interface etc
func NewLexer(input string) *lexer {
	return &lexer{
		rawJSON:   input,
		index:     0,
		nextIndex: 1,
	}
}

func (l *lexer) NextToken() token.Token {
	// Return EOF token when index reached out of range in JSON length
	if len(l.rawJSON) <= l.index {
		return token.Token{Type: token.EOF}
	}

	ch := l.rawJSON[l.index]
	tok := token.Token{Type: token.ILLEGAL}

	switch ch {
	case '{':
		tok = token.Token{
			Type:    token.LBRACE,
			Literal: string(ch),
		}
	case '}':
		tok = token.Token{
			Type:    token.RBRACE,
			Literal: string(ch),
		}
	case ':':
		panic("case : : not implement")
	case ',':
		panic("case , : not implement")
	default:
		{
			// string literal
		}
		{
			// number case
		}
		{
			// reserved word
		}
	}

	// Proceed Lexer index
	l.index += 1
	l.nextIndex = l.index + 1

	return tok
}
