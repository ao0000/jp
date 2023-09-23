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
		tok = token.Token{
			Type:    token.COLON,
			Literal: string(ch),
		}
	case ',':
		tok = token.Token{
			Type:    token.COMMA,
			Literal: string(ch),
		}
	case '[':
		tok = token.Token{
			Type:    token.LBRACK,
			Literal: string(ch),
		}
	case ']':
		tok = token.Token{
			Type:    token.RBRACK,
			Literal: string(ch),
		}
	default:
		// String literal
		if ch == '"' {
			ltr := ""
			for {
				l.index += 1
				l.nextIndex = l.index + 1

				// TODO: Implement to escape Double Quotation(\")
				if l.rawJSON[l.index] == '"' {
					tok = token.Token{
						Type:    token.STRING,
						Literal: ltr,
					}
					break
				}
				ltr += string(l.rawJSON[l.index])
			}
		} else if isLetter(ch) {
			idnt := string(ch)
			for {
				nextCh := l.rawJSON[l.nextIndex]
				if isLetter(nextCh) {
					idnt += string(nextCh)
				} else {
					break
				}

				l.index += 1
				l.nextIndex = l.index + 1
			}

			if tokType, ok := token.ReservedWords[idnt]; ok {
				tok = token.Token{
					Type:    tokType,
					Literal: idnt,
				}
			} else {
				tok = token.Token{
					Type:    token.ILLEGAL,
					Literal: idnt,
				}
			}
		}
		{
			// number case
		}
	}

	// Proceed Lexer index
	l.index += 1
	l.nextIndex = l.index + 1

	return tok
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
