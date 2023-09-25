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
		{"Input JSON argument is a object has a set of key and string value",
			"{\"key\":\"value\"}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.STRING, Literal: "value"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has some sets of key and string value",
			"{\"key1\":\"value1\",\"key2\":\"value2\"}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key1"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.STRING, Literal: "value1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.STRING, Literal: "key2"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.STRING, Literal: "value2"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a sets of key and reserved words",
			"{\"key1\":null,\"key2\":true,\"key3\":false}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key1"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NULL, Literal: "null"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.STRING, Literal: "key2"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.STRING, Literal: "key3"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.FALSE, Literal: "false"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a set of key and null and an array",
			"{\"key1\":null,\"key2\":[true,false]}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key1"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NULL, Literal: "null"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.STRING, Literal: "key2"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACK, Literal: "["},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.FALSE, Literal: "false"},
				{Type: token.RBRACK, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a set of key and number value",
			"{\"key\":1}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: "1"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a set of key and negative number value",
			"{\"key\":-1}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: "-1"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a set of key and decimal number value",
			"{\"key\":1.1}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.NUMBER, Literal: "1.1"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a set of key and illegal value",
			"{\"key\":illegal}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.ILLEGAL, Literal: "illegal"},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is a object has a set of key and illegal character value",
			"{\"key\":*}",
			[]token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "key"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.ILLEGAL},
				{Type: token.RBRACE, Literal: "}"},
			},
		},
		{"Input JSON argument is empty object including white space",
			"{ }",
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
