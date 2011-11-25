package lex

import "testing"

func Test_EmptyShouldReturnEOF(t *testing.T){
	_, ch := NewLexer("")
	if tok := <-ch; tok.Value != eof {
		t.Error("An empty lexer should only return EOF")
	}
}