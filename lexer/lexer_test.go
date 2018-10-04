package lexer

import (
	"fmt"
	"testing"

	"github.com/qq52184962/compiler/token"
)

func Test_readString(t *testing.T) {
	strs := []string{"abcde", " a  c e ", "asd f", "12sd ", "hi 你好！"}
	for i, s := range strs {
		str := fmt.Sprintf(`"%s"`, s)
		lexer := New()
		lexer.input = []byte(str)
		tok := lexer.nextToken()
		if tok.Type != token.STRINGLITERAL || tok.Data != s {
			t.Errorf("%d readString is wrong, expected to be %s, but get %s", i, s, tok.Data)
		}
	}
}
