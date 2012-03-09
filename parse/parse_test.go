package parse

import (
	"github.com/maxpolun/tundra/lex"
	"testing"
)

func Test_ShouldReturnNilJustEOF(t *testing.T) {
	_, ch := lex.NewLexer("")
	p := NewParser(ch)
	if tree := p.Parse(); tree != nil {
		t.Errorf("Parse on empty string should only return nil, got %v instead", tree)
	}
}

func Test_IntegerShouldReturnIntegerLiterals(t *testing.T) {
	_, ch := lex.NewLexer("123")
	p := NewParser(ch)
	if tree := p.Parse(); tree.Eval() != 123 {
		t.Errorf("Parse Expected to produce %v, got %v", 123, tree.Eval())
	}
}

func makeFakeChan(tokens []lex.Token) (l *lex.Lexer, ch chan lex.Token) {
	ch = make(chan lex.Token)
	go func() {
		for _, t := range tokens {
			ch <- t
		}
		close(ch)
	}()
	return nil, ch
}

func Test_PlusMinusShouldParse(t *testing.T) {
	tokens := [][]lex.Token{
		[]lex.Token{
			lex.Token{lex.T_INT, "1"}, lex.Token{lex.T_PLUS, "+"}, lex.Token{lex.T_INT, "1"}}}

	expected := []int{
		2}
	for index := range tokens {
		_, ch := makeFakeChan(tokens[index])
		p := NewParser(ch)
		if tree := p.Parse(); tree.Eval() != expected[index] {
			t.Errorf("Eval() produced %v, expected %v", tree.Eval(), expected[index])
		}
	}
}
func Test_PlusShouldEval(t *testing.T) {
	tree := &BinOp{
		Left:  IntLiteral(1),
		Right: IntLiteral(1),
		Op: func(a, b int) int {
			return a + b
		},
		text: "+"}
	result := tree.Eval()
	if result != 2 {
		t.Errorf("%v.Eval() produced %v should produce 2", tree, result)
	}
}
