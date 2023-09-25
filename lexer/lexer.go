package lexer

import (
	"github.com/ao0000/jp/token"
)

type Lexer interface {
	NextToken() token.Token
}

type lexer struct {
	rawJSON   string
	index     int
	nextIndex int
}

func NewLexer(input string) Lexer {
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
	var tok token.Token

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
		if ch == '"' {
			// String literal
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
			// Reserve words
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
		} else if isNumber(ch) || ch == '-' {
			// number case
			num := string(ch)
			for {
				nextCh := l.rawJSON[l.nextIndex]
				if isNumber(nextCh) || nextCh == '.' {
					num += string(nextCh)
				} else {
					break
				}
				l.index += 1
				l.nextIndex = l.index + 1
			}
			tok = token.Token{
				Type:    token.NUMBER,
				Literal: num,
			}
		} else {
			tok = token.Token{Type: token.ILLEGAL}
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

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
