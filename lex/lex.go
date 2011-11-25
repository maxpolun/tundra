package lex

import (
//	"io"
//	"unicode"
)

const (
	T_INT =  iota
	T_PLUS
	T_MINUS
	T_LPAREN
	T_RPAREN
)


const eof = -1

type Token struct {
	Value int
	Text  string
}

var tokenValues = map[int]string {
	T_INT: "T_INT",
	T_PLUS: "T_PLUS",
	T_MINUS: "T_MINUS",
	T_LPAREN: "T_LPAREN",
	T_RPAREN: "T_RPAREN",
}

func (t Token) String() string {
	return t.Text
}

type Lexer struct {
	input string
	start int
	pos   int
	width int

	tokens chan Token
}

type stateFn func(*Lexer) stateFn

func NewLexer(s string) (*Lexer, chan Token) {
	l := &Lexer{
		
	}
	l.input = s
	l.tokens = make(chan Token)
	go l.run()

	return l, l.tokens

}

func (l *Lexer) run() {
	/*
	for state := defaultState; state != nil; {
		state = state(l)
	}
	*/
	l.tokens <- Token{Value:eof}
	close(l.tokens)
}
func emit(l *Lexer, value int) {
	l.tokens <- Token{Value: value, Text:l.input[l.start:l.pos]}
	l.start = l.pos
}
func (l *Lexer) next() (rune int) {
	if (l.pos > len(l.input)) {
		l.width = 0
		return eof;
	}
	return 0
}
func defaultState(l *Lexer) stateFn {
	return defaultState
}








