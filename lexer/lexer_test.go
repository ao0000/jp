package lexer_test

import (
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		name string
		JSON string
		want []token.Token
	}{
		{"Input JSON argument is empty object",
			"{}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.JSON)
			tokens := []token.Token{}

			for {
				tok := l.NextToken()
				if tok.Type == token.EOF {
					break
				}

				tokens = append(tokens, tok)
			}

			for i, tok := range tokens {
				if tok != tt.want[i] {
					t.Fatalf("tokens: %+v, want:%+v", tokens, tt.want)
				}
			}
		})
	}
}
