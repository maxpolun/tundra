package lex

import "testing"

func Test_EmptyShouldReturnEOF(t *testing.T) {
	_, ch := NewLexer("")
	if tok := <-ch; tok.Value != eof {
		t.Errorf("An empty lexer should only return EOF, got %v instead", tok.Value)
	}
}

func Test_SymbolShouldReturnTokenValue(t *testing.T) {
	syms := []string{
		"+",
		"-",
		"(",
		")"}
	expected := []int{
		T_PLUS,
		T_MINUS,
		T_LPAREN,
		T_RPAREN}
	for i := range syms {
		_, ch := NewLexer(syms[i])
		if tok := <-ch; tok.Value != expected[i] {
			t.Errorf("for string %v, expected %v, got %v", syms[i], expected[i], tok.Value)
		}
	}

}
func Test_MultipleSymbolsShouldEmitMultipleTokens(t *testing.T) {
	syms := "+-()"
	expected := []int{
		T_PLUS,
		T_MINUS,
		T_LPAREN,
		T_RPAREN,
		eof}

	_, ch := NewLexer(syms)
	i := 0
	for tok := range ch {
		if tok.Value != expected[i] {
			t.Errorf("for string %c, expected %v, got %v", syms[i], expected[i], tok.Value)
		}
		i++
	}
}

func Test_ShouldLexInts(t *testing.T) {
	text := "123"
	expected := T_INT

	_, ch := NewLexer(text)
	if tok := <-ch; tok.Value != expected || tok.Text != text {
		t.Errorf("for string %v got %v(%v), expected %v", text, tok.Value, tok.Text, expected)
	}
}
