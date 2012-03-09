package parse

import (
	"fmt"
	"github.com/maxpolun/tundra/lex"
	"strconv"
	"strings"
)

type Parser struct {
	ch chan lex.Token
}

type ASTNode interface {
	Eval() int
}
type IntLiteral int

func (i IntLiteral) Eval() int {
	return int(i)
}
func (i IntLiteral) String() string {
	return fmt.Sprint(int(i))
}

type BinOp struct {
	Left  ASTNode
	Right ASTNode
	Op    func(int, int) int
	text  string
}

func (b *BinOp) Eval() int {
	return b.Op(b.Left.Eval(), b.Right.Eval())
}
func (b *BinOp) String() string {
	return fmt.Sprintf("%v(%v, %v)", b.text, b.Left, b.Right)
}

func NewParser(c chan lex.Token) Parser {
	p := Parser{ch: c}
	return p
}

var binopTable = map[string]func(int, int) int{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	}}

func parseLit(tok lex.Token) ASTNode {
	switch tok.Value {
	case lex.T_INT:
		num, _ := strconv.Atoi(tok.Text)
		return IntLiteral(num)
	default:
		return nil
	}
	return nil
}

func strTokList(tokList []lex.Token) string {
	stringList := make([]string, len(tokList))
	for i := range tokList {
		stringList[i] = tokList[i].Text
	}
	return strings.Join(stringList, ", ")
}

func (p *Parser) Parse() ASTNode {

	tokList := make([]lex.Token, 0, 10)

	for tok := range p.ch {
		tokList = append(tokList, tok)
	}
	fmt.Printf("number of tokens to parse: %v\n", len(tokList))
	fmt.Printf("tokens: %v\n", strTokList(tokList))
	if len(tokList) == 0 {
		return nil
	}
	if len(tokList) == 1 {
		return parseLit(tokList[0])
	}
	if tokList[1].Value == lex.T_PLUS {
		return &BinOp{Left: parseLit(tokList[0]),
			Right: parseLit(tokList[2]),
			Op:    binopTable["+"],
			text:  tokList[1].Text}
	} else {
		return parseLit(tokList[0])
	}
	return nil
}
