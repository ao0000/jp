package ast

import (
	"fmt"
	"github.com/ao0000/jp/token"
)

type Value interface {
	Literal() string
	String() string
	Value() Value
}

type JSON struct {
	value Value
}

var _ Value = (*JSON)(nil)

func NewJSON(value Value) *JSON {
	return &JSON{value: value}
}

func (j *JSON) Literal() string {
	return j.value.Literal()
}

func (j *JSON) String() string {
	return j.value.String()
}

func (j *JSON) Value() Value {
	return j.value
}

type Object struct {
	tok    token.Token
	Keys   []*String
	Values []Value
}

var _ Value = (*Object)(nil)

func NewObject(keys []*String, values []Value) *Object {
	if len(keys) == 0 && len(values) == 0 {
		return &Object{tok: token.Token{Type: token.OBJECT, Literal: "{}"}, Keys: keys, Values: values}
	}

	literal := "{"
	reg := len(keys)
	if reg > len(values) {
		reg = len(values)
	}
	for i := 0; i < reg; i++ {
		if literal != "{" {
			literal += ","
		}
		literal += fmt.Sprintf("%s:%s", keys[i], values[i])
	}
	literal += "}"

	tok := token.Token{Type: token.OBJECT, Literal: literal}
	return &Object{tok: tok, Keys: keys, Values: values}
}

func (o *Object) Literal() string {
	return o.tok.Literal
}

func (o *Object) String() string {
	return o.tok.Literal
}

func (o *Object) Value() Value {
	return o
}

type Array struct {
	tok    token.Token
	Values []Value
}

var _ Value = (*Array)(nil)

func NewArray(value []Value) *Array {
	lit := fmt.Sprint("%+v", value)
	return &Array{tok: token.Token{Type: token.ARRAY, Literal: lit}, Values: value}
}

func (a *Array) Literal() string {
	return a.tok.Literal
}

func (a *Array) String() string {
	return a.tok.Literal
}

func (a *Array) Value() Value {
	return a
}

type Boolean struct {
	tok   token.Token
	value bool
}

var _ Value = (*Boolean)(nil)

func NewBoolean(tok token.Token) (*Boolean, error) {
	if tok.Type == token.TRUE {
		return &Boolean{tok: tok, value: true}, nil

	} else if tok.Type == token.FALSE {
		return &Boolean{tok: tok, value: false}, nil
	}
	return nil, fmt.Errorf("failed to parse enexpected bolean token: %+v", tok)
}

func (b *Boolean) Literal() string {
	return b.tok.Literal
}

func (b *Boolean) String() string {
	return b.tok.Literal
}

func (b *Boolean) Value() Value {
	return b
}

type String struct {
	tok   token.Token
	value string
}

var _ Value = (*String)(nil)

func NewString(tok token.Token) *String {
	return &String{tok: tok, value: tok.Literal}
}

func (s *String) Literal() string {
	return fmt.Sprintf("\"%s\"", s.tok.Literal)
}

func (s *String) String() string {
	return s.tok.Literal
}

func (s *String) Value() Value {
	return s
}

type Number[T int64 | float64] struct {
	tok   token.Token
	value T
}

var _ Value = (*Number[int64])(nil)

func NewNumber[T int64 | float64](tok token.Token, value T) *Number[T] {
	return &Number[T]{tok: tok, value: value}
}

func (n *Number[T]) Literal() string {
	return n.tok.Literal
}

func (n *Number[T]) String() string {
	return n.tok.Literal
}

func (n *Number[T]) Value() Value {
	return n
}

type Null struct {
	tok   token.Token
	value interface{}
}

var _ Value = (*Null)(nil)

func NewNull(tok token.Token) *Null {
	return &Null{tok: tok, value: nil}
}

func (n *Null) Literal() string {
	return n.tok.Literal
}

func (n *Null) String() string {
	return n.tok.Literal
}

func (n *Null) Value() Value {
	return n
}
