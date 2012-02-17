package parse

import (
	"../lex"
)

func Main(str string) {
	l, ch := lex.NewLexer(str)
	for tok := range ch {

	}
}
