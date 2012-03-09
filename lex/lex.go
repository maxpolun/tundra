package lex

import (
	//"fmt"
	"unicode"
	"unicode/utf8"
)

const (
	T_INT = iota
	T_PLUS
	T_MINUS
	T_LPAREN
	T_RPAREN
)

const EOF = -1

type Token struct {
	Value int
	Text  string
}

var tokenValues = map[int]string{
	T_INT:    "T_INT",
	T_PLUS:   "T_PLUS",
	T_MINUS:  "T_MINUS",
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
	l := &Lexer{}
	l.input = s
	l.tokens = make(chan Token)
	l.pos = 0
	go l.run()

	return l, l.tokens

}

func (l *Lexer) run() {
	/*
		for state := defaultState; state != nil; {
			state = state(l)
		}
	*/

	for r := l.next(); r != EOF; r = l.next() {
		l.emitSingle(r)
	}
	l.emit(EOF)
	close(l.tokens)
}

func (l *Lexer) lexInt() {
	for r := l.next(); unicode.IsDigit(r) && r != EOF; r = l.next() {
	}
	l.backup()
	l.emit(T_INT)
}

func (l *Lexer) emitSingle(r rune) bool {
	if unicode.IsDigit(r) {
		l.backup()
		l.lexInt()
		return true
	}
	switch r {
	case '+':
		l.emit(T_PLUS)
		return true
	case '-':
		l.emit(T_MINUS)
		return true
	case '(':
		l.emit(T_LPAREN)
		return true
	case ')':
		l.emit(T_RPAREN)
		return true
	default:

	}
	return false
}

func (l *Lexer) emit(value int) {
	l.tokens <- Token{Value: value, Text: l.input[l.start:l.pos]}
	l.start = l.pos
}
func (l *Lexer) next() (r rune) {
	if l.pos >= len(l.input) {

		l.width = 0
		return EOF
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}
func (l *Lexer) backup() {
	l.pos -= l.width
}
func defaultState(l *Lexer) stateFn {
	return defaultState
}
