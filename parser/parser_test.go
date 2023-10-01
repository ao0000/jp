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
				ast.NewObject([]*ast.String{}, []ast.Value{}),
			),
		},
		{
			name: "object string and string case",
			str:  "{\"key\":\"value\"}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{ast.NewString(token.Token{token.STRING, "value"})},
				),
			),
		},
		{
			name: "object string and number case",
			str:  "{\"key\":123}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{ast.NewNumber[int64](token.Token{token.NUMBER, "123"}, 123)},
				),
			),
		},
		{
			name: "object string and negative number case",
			str:  "{\"key\":-123}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{ast.NewNumber[int64](token.Token{token.NUMBER, "-123"}, -123)},
				),
			),
		},
		{
			name: "object string and float number case",
			str:  "{\"key\":1.23}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{ast.NewNumber[float64](token.Token{token.NUMBER, "1.23"}, 1.23)},
				),
			),
		},
		{
			name: "object string and boolean case",
			str:  "{\"key\":true}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{newTrueToken()},
				),
			),
		},
		{
			name: "object string and null case",
			str:  "{\"key\":null}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{ast.NewNull(token.Token{token.NULL, "null"})},
				),
			),
		},
		{
			name: "object string and nest object case",
			str:  "{\"key\":{\"key1\":\"value1\"}}",
			want: ast.NewJSON(
				ast.NewObject(
					[]*ast.String{ast.NewString(token.Token{token.STRING, "key"})},
					[]ast.Value{
						ast.NewObject(
							[]*ast.String{ast.NewString(token.Token{token.STRING, "key1"})},
							[]ast.Value{ast.NewString(token.Token{token.STRING, "value1"})},
						),
					},
				),
			),
		},
		{
			name: "array string case",
			str:  "[]",
			want: ast.NewJSON(
				ast.NewArray([]ast.Value{}),
			),
		},
		{
			name: "array string case",
			str:  "[\"value1\"]",
			want: ast.NewJSON(
				ast.NewArray([]ast.Value{
					ast.NewString(token.Token{token.STRING, "value1"}),
				}),
			),
		},
		{
			name: "array string case",
			str:  "[\"value1\",\"value2\"]",
			want: ast.NewJSON(
				ast.NewArray([]ast.Value{
					ast.NewString(token.Token{token.STRING, "value1"}),
					ast.NewString(token.Token{token.STRING, "value2"}),
				}),
			),
		},
		{
			name: "array string case",
			str:  "[{\"key1\",\"value1\"},{\"key2\",\"value2\"}]",
			want: ast.NewJSON(
				ast.NewArray([]ast.Value{
					ast.NewObject(
						[]*ast.String{ast.NewString(token.Token{token.STRING, "key1"})},
						[]ast.Value{ast.NewString(token.Token{token.STRING, "value1"})},
					),
					ast.NewObject(
						[]*ast.String{ast.NewString(token.Token{token.STRING, "key2"})},
						[]ast.Value{ast.NewString(token.Token{token.STRING, "value2"})},
					),
				}),
			),
		},
		{
			name: "array number case",
			str:  "[123,234]",
			want: ast.NewJSON(
				ast.NewArray(
					[]ast.Value{
						ast.NewNumber[int64](token.Token{token.NUMBER, "123"}, 123),
						ast.NewNumber[int64](token.Token{token.NUMBER, "234"}, 234),
					}),
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
