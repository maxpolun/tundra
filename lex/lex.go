package lex

import (
	//"fmt"
	"unicode"
	"utf8"
)

const (
	T_INT = iota
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

	for rune := l.next(); rune != eof; rune = l.next() {
		l.emitSingle(rune)
	}
	l.emit(eof)
	close(l.tokens)
}

func (l *Lexer) endOfFile() {
	l.tokens <- Token{Value: eof}
}

func (l *Lexer) lexInt() {
	for rune := l.next(); unicode.IsDigit(rune); rune = l.next() {
	}
	l.backup()
	l.emit(T_INT)
}

func (l *Lexer) emitSingle(rune int) bool {
	if unicode.IsDigit(rune) {
		l.backup()
		l.lexInt()
	}
	switch rune {
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
func (l *Lexer) next() (rune int) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}
func (l *Lexer) backup() {
	l.pos -= l.width
}
func defaultState(l *Lexer) stateFn {
	return defaultState
}
