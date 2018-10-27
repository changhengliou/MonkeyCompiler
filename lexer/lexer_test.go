package lexer

import (
	"fmt"
	"testing"

	"github.com/qq52184962/MonkeyCompiler/token"
)

func Test_readString(t *testing.T) {
	strs := []string{"abcde", " a  c e ", "asd f", "12sd ", "hi 你好！"}
	for i, s := range strs {
		str := fmt.Sprintf(`"%s"`, s)
		lexer := New()
		lexer.input = []byte(str)
		tok := lexer.NextToken()
		if tok.Type != token.STRINGLITERAL || tok.Data != s {
			t.Errorf("%d readString is wrong, expected to be %s, but get %s", i, s, tok.Data)
		}
	}
}

func Test_readIdentifier(t *testing.T) {
	strs := []string{"let x = 9;",
		`let variable = "word"`,
		"let w3cf9rk03 = 1",
	}
	ans := []string{"x", "variable", "w3cf9rk03"}
	for i, str := range strs {
		lexer := New()
		lexer.input = []byte(str)
		lexer.currPos = 3
		tok := lexer.NextToken()
		if tok.Type != token.IDENTIFIER || tok.Data != ans[i] {
			t.Errorf("[%d] readIdentifier: expected to be %s, but get %s (%s)", i, ans[i], tok.Data, tok.Type)
		}
	}
}

func Test_nextToken(t *testing.T) {
	strs := []string{
		`let a = 1;`,
		`let b = "text";`,
		`let c38f6_sa = 3 * (23 + 1);`,
		`let __arr__ = [ 1, 2, 3 ];`,
		`let obj = { w: 1, x: 2 };`,
		`obj["w"] != 1;`,
		`let add = func(x, y) { x + y; };
		 add(1, 2);`,
	}
	ans := [][]token.Token{
		{
			token.Token{Type: token.LET, Data: "let"},
			token.Token{Type: token.IDENTIFIER, Data: "a"},
			token.Token{Type: token.ASSIGN, Data: byte('=')},
			token.Token{Type: token.NUMBER, Data: 1.0},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
		{
			token.Token{Type: token.LET, Data: "let"},
			token.Token{Type: token.IDENTIFIER, Data: "b"},
			token.Token{Type: token.ASSIGN, Data: byte('=')},
			token.Token{Type: token.STRINGLITERAL, Data: "text"},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
		{
			token.Token{Type: token.LET, Data: "let"},
			token.Token{Type: token.IDENTIFIER, Data: "c38f6_sa"},
			token.Token{Type: token.ASSIGN, Data: byte('=')},
			token.Token{Type: token.NUMBER, Data: 3.0},
			token.Token{Type: token.MULTIPLY, Data: byte('*')},
			token.Token{Type: token.LPAREN, Data: byte('(')},
			token.Token{Type: token.NUMBER, Data: 23.0},
			token.Token{Type: token.PLUS, Data: byte('+')},
			token.Token{Type: token.NUMBER, Data: 1.0},
			token.Token{Type: token.RPAREN, Data: byte(')')},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
		{ // let __arr__ = [ 1, 2, 3 ];
			token.Token{Type: token.LET, Data: "let"},
			token.Token{Type: token.IDENTIFIER, Data: "__arr__"},
			token.Token{Type: token.ASSIGN, Data: byte('=')},
			token.Token{Type: token.LSQRBRACKET, Data: byte('[')},
			token.Token{Type: token.NUMBER, Data: 1.0},
			token.Token{Type: token.COMMA, Data: byte(',')},
			token.Token{Type: token.NUMBER, Data: 2.0},
			token.Token{Type: token.COMMA, Data: byte(',')},
			token.Token{Type: token.NUMBER, Data: 3.0},
			token.Token{Type: token.RSQRBRACKET, Data: byte(']')},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
		{ // let obj = { w: 1, x: 2 };
			token.Token{Type: token.LET, Data: "let"},
			token.Token{Type: token.IDENTIFIER, Data: "obj"},
			token.Token{Type: token.ASSIGN, Data: byte('=')},
			token.Token{Type: token.LBRACKET, Data: byte('{')},
			token.Token{Type: token.IDENTIFIER, Data: "w"},
			token.Token{Type: token.COLON, Data: byte(':')},
			token.Token{Type: token.NUMBER, Data: 1.0},
			token.Token{Type: token.COMMA, Data: byte(',')},
			token.Token{Type: token.IDENTIFIER, Data: "x"},
			token.Token{Type: token.COLON, Data: byte(':')},
			token.Token{Type: token.NUMBER, Data: 2.0},
			token.Token{Type: token.RBRACKET, Data: byte('}')},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
		{ // obj["w"] != 1;
			token.Token{Type: token.IDENTIFIER, Data: "obj"},
			token.Token{Type: token.LSQRBRACKET, Data: byte('[')},
			token.Token{Type: token.STRINGLITERAL, Data: "w"},
			token.Token{Type: token.RSQRBRACKET, Data: byte(']')},
			token.Token{Type: token.NOTEQUAL, Data: "!="},
			token.Token{Type: token.NUMBER, Data: 1.0},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
		{ // let add = func(x, y) { x + y; };
			// add(1, 2);
			token.Token{Type: token.LET, Data: "let"},
			token.Token{Type: token.IDENTIFIER, Data: "add"},
			token.Token{Type: token.ASSIGN, Data: byte('=')},
			token.Token{Type: token.FUNC, Data: "func"},
			token.Token{Type: token.LPAREN, Data: byte('(')},
			token.Token{Type: token.IDENTIFIER, Data: "x"},
			token.Token{Type: token.COMMA, Data: byte(',')},
			token.Token{Type: token.IDENTIFIER, Data: "y"},
			token.Token{Type: token.RPAREN, Data: byte(')')},
			token.Token{Type: token.LBRACKET, Data: byte('{')},
			token.Token{Type: token.IDENTIFIER, Data: "x"},
			token.Token{Type: token.PLUS, Data: byte('+')},
			token.Token{Type: token.IDENTIFIER, Data: "y"},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
			token.Token{Type: token.RBRACKET, Data: byte('}')},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
			token.Token{Type: token.EOL, Data: byte('\n')},
			token.Token{Type: token.IDENTIFIER, Data: "add"},
			token.Token{Type: token.LPAREN, Data: byte('(')},
			token.Token{Type: token.NUMBER, Data: 1.0},
			token.Token{Type: token.COMMA, Data: byte(',')},
			token.Token{Type: token.NUMBER, Data: 2.0},
			token.Token{Type: token.RPAREN, Data: byte(')')},
			token.Token{Type: token.SEMICOLON, Data: byte(';')},
		},
	}
	for i, s := range strs {
		lexer := New()
		lexer.input = []byte(s)
		j := 0
		for lexer.currPos < len(lexer.input) && j < len(ans[i]) {
			tok := lexer.NextToken()
			if tok.Data != ans[i][j].Data || tok.Type != ans[i][j].Type {
				t.Errorf("[%d] nextToken: expected to be %s (%s), but get %s (%s)",
					i, ans[i][j].Data, ans[i][j].Type, tok.Data, tok.Type)
			}
			j++
		}
		if j != len(ans[i]) {
			t.Errorf("[%d] nextToken: expected to get %d tokens, but only %d are return", i, len(ans[i]), j)
		}
	}
}
