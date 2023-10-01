package parser_test

import (
	"github.com/ao0000/jp/ast"
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/parser"
	"github.com/ao0000/jp/token"
	"reflect"
	"testing"
)

func Test_parser_Parse(t *testing.T) {
	// TODO: refactor test case name
	tests := []struct {
		name string
		str  string
		want *ast.JSON
	}{
		{
			name: "empty object case",
			str:  "{}",
			want: ast.NewJSON(
				ast.NewObject(map[*ast.String]ast.Value{}),
			),
		},
		{
			name: "object string and string case",
			str:  "{\"key\":\"value\"}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): ast.NewString(token.Token{token.STRING, "value"}),
					},
				),
			),
		},
		{
			name: "object string and number case",
			str:  "{\"key\":123}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): ast.NewNumber[int64](token.Token{token.NUMBER, "123"}, 123),
					},
				),
			),
		},
		{
			name: "object string and negative number case",
			str:  "{\"key\":-123}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): ast.NewNumber[int64](token.Token{token.NUMBER, "-123"}, -123),
					},
				),
			),
		},
		{
			name: "object string and float number case",
			str:  "{\"key\":1.23}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): ast.NewNumber[float64](token.Token{token.NUMBER, "1.23"}, 1.23),
					},
				),
			),
		},
		{
			name: "object string and boolean case",
			str:  "{\"key\":true}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): newTrueToken(),
					},
				),
			),
		},
		{
			name: "object string and null case",
			str:  "{\"key\":null}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): ast.NewNull(token.Token{token.NULL, "null"}),
					},
				),
			),
		},
		{
			name: "object string and nest object case",
			str:  "{\"key\":{\"key1\":\"value1\"}}",
			want: ast.NewJSON(
				ast.NewObject(
					map[*ast.String]ast.Value{
						ast.NewString(token.Token{token.STRING, "key"}): ast.NewObject(
							map[*ast.String]ast.Value{
								ast.NewString(token.Token{token.STRING, "key1"}): ast.NewString(token.Token{token.STRING, "value1"}),
							},
						),
					},
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer(tt.str)
			p := parser.NewParser(l)
			ast := p.Parse()
			if !reflect.DeepEqual(ast, tt.want) {
				t.Fatalf("ast: %+v, want:%+v", ast, tt.want)
			}
		})
	}
}

func newTrueToken() *ast.Boolean {
	tok, _ := ast.NewBoolean(token.Token{token.TRUE, "true"})
	return tok
}
