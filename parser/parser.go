package parser

import (
	"fmt"
	"github.com/ao0000/jp/ast"
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/token"
	"strconv"
)

type parser struct {
	l         lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []error
}

func NewParser(lxr lexer.Lexer) *parser {
	p := &parser{l: lxr, errors: []error{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *parser) Errors() []error {
	return p.errors
}

func (p *parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *parser) Parse() *ast.JSON {
	json := ast.NewJSON(p.parseJSON())
	return json
}

func (p *parser) parseJSON() ast.Value {
	switch tok := p.curToken.Type; tok {
	case token.LBRACE:
		// Object
		return p.parseObject()
	case token.LBRACK:
		// Array
		return p.parseArray()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token: %+v", tok))
	}
	return nil
}

func (p *parser) parseObject() *ast.Object {
	p.nextToken()
	if p.curToken.Type == token.RBRACE {
		return ast.NewObject(nil, nil)
	}
	var key *ast.String
	var value ast.Value

	if p.curToken.Type != token.STRING {
		p.errors = append(p.errors, fmt.Errorf("object key is unexpected token: %+v", p.curToken))
	}
	key = p.parseString()
	if p.curToken.Type != token.COLON {
		p.errors = append(p.errors, fmt.Errorf("infix between object key and value unexpected token: %+v", p.curToken))
	}
	p.nextToken()

	value = p.parseValue()
	if p.curToken.Type != token.RBRACE {
		p.errors = append(p.errors, fmt.Errorf("object right brace is unexpected token: %+v", p.curToken))
	}
	p.nextToken()

	return ast.NewObject(key, value)
}

func (p *parser) parseArray() *ast.Array {
	p.nextToken()
	v := []ast.Value{}
	for {
		if p.curToken.Type == token.RBRACK {
			p.nextToken()
			ast.NewArray(v)
		}
		p.parseValue()
		if p.curToken.Type == token.COMMA {
			p.nextToken()
		} else {
			p.errors = append(p.errors, fmt.Errorf("failed to parseArray unexpected token: %+v", p.curToken))
			p.nextToken()
		}
	}
}

func (p *parser) parseValue() ast.Value {
	tok := p.curToken
	typ := tok.Type
	switch {
	case typ == token.NULL:
		return p.parseNull()
	case typ == token.TRUE || typ == token.FALSE:
		return p.parseBoolean()
	case typ == token.STRING:
		// string
		return p.parseString()
	case typ == token.NUMBER:
		// number
		for _, ch := range tok.Literal {
			if ch == '.' {
				return p.parseFloat()
			}
		}
		return p.parseInteger()

	case typ == token.LBRACE:
		// object
		return p.parseObject()
	case typ == token.LBRACK:
		// array
		return p.parseArray()
	}
	p.errors = append(p.errors, fmt.Errorf("failed to parseValue"))
	return nil
}

func (p *parser) parseNull() *ast.Null {
	tok := p.curToken
	if tok.Type != token.NULL {
		p.errors = append(p.errors, fmt.Errorf("null is unpexpected token: %+v", p.curToken))
	}
	p.nextToken()
	return ast.NewNull(tok)
}

func (p *parser) parseBoolean() *ast.Boolean {
	b, err := ast.NewBoolean(p.curToken)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("failed to parse enexpected bolean token: %+v", p.curToken))
	}
	p.nextToken()
	return b
}

func (p *parser) parseString() *ast.String {
	tok := p.curToken
	p.nextToken()
	return ast.NewString(tok)
}

func (p *parser) parseInteger() *ast.Number[int64] {
	tok := p.curToken
	p.nextToken()
	v, err := strconv.ParseInt(tok.Literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("failed to parseInteger: %w", err))
	}
	return ast.NewNumber(tok, v)
}

func (p *parser) parseFloat() *ast.Number[float64] {
	tok := p.curToken
	p.nextToken()
	v, err := strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("failed to parseFloat: %w", err))
	}
	return ast.NewNumber(tok, v)
}
