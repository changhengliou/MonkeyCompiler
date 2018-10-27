package lexer

import (
	"fmt"
	"github.com/qq52184962/MonkeyCompiler/token"
)

type Lexer struct {
	input      []byte
	currPos    int
	lineNumber int
	position   int
}

func New() *Lexer {
	return &Lexer{currPos: 0, lineNumber: 0, position: 0}
}

func (lex *Lexer) NextToken() token.Token {
	var t token.Token
	if lex.currPos >= len(lex.input) {
		t.Type = token.EOF
	} else {
		// skip whitespace
		for lex.input[lex.currPos] == ' ' || lex.input[lex.currPos] == '\t' {
			lex.currPos++
		}

		switch lex.input[lex.currPos] {
		case '\n':
			t.Type = token.EOL
			break
		case '(':
			t.Type = token.LPAREN
			break
		case ')':
			t.Type = token.RPAREN
			break
		case '{':
			t.Type = token.LBRACKET
			break
		case '}':
			t.Type = token.RBRACKET
			break
		case ',':
			t.Type = token.COMMA
			break
		case ';':
			t.Type = token.SEMICOLON
			break
		case ':':
			t.Type = token.COLON
			break
		case '"':
			lex.currPos++
			t.Type = token.STRINGLITERAL
			t.Data = lex.readString()
			break
		case '=':
			if lex.input[lex.currPos+1] == '=' {
				t.Type = token.EQUAL
				t.Data = "=="
				lex.currPos++
			} else {
				t.Type = token.ASSIGN
			}
			break
		case '!':
			if lex.input[lex.currPos+1] == '=' {
				t.Type = token.NOTEQUAL
				t.Data = "!="
				lex.currPos++
			} else {
				t.Type = token.NOT
			}
			break
		case '+':
			t.Type = token.PLUS
			break
		case '-':
			t.Type = token.MINUS
			break
		case '*':
			t.Type = token.MULTIPLY
			break
		case '/':
			t.Type = token.DIVIDE
			break
		case '>':
			t.Type = token.GREATER
			break
		case '<':
			t.Type = token.SMALLER
			break
		case '[':
			t.Type = token.LSQRBRACKET
			break
		case ']':
			t.Type = token.RSQRBRACKET
			break
		case 0:
			t.Type = token.EOF
			break
		default:
			if isDigit(lex.input[lex.currPos]) {
				t.Type = token.NUMBER
				t.Data = lex.readNumber()
			} else {
				t.Data = lex.readIdentifier()
				t.Type = token.GetKeywordOrIdentifier(t.Data.(string))
			}
			break
		}
		if t.Data == nil {
			t.Data = lex.input[lex.currPos]
		}
		lex.currPos++
	}
	return t
}

func (lex *Lexer) readNumber() float64 {
	var f float64
	f = float64(lex.readInt())
	if lex.input[lex.currPos] == '.' {
		f += lex.readDecimal()
	}
	return f
}

func (lex *Lexer) readDecimal() float64 {
	var f float64
	for i := 1.0; lex.currPos < len(lex.input) && isDigit(lex.input[lex.currPos]); lex.currPos++ {
		f = i / 10.0 * float64(lex.input[lex.currPos]-'0')
	}
	return f
}

func (lex *Lexer) readInt() int {
	var f int
	for ; lex.currPos < len(lex.input) && isDigit(lex.input[lex.currPos]); lex.currPos++ {
		f = f*10 + int(lex.input[lex.currPos]-'0')
	}
	lex.currPos--
	return f
}

func (lex *Lexer) readString() string {
	str := make([]byte, 0)
	for ; lex.currPos < len(lex.input) && lex.input[lex.currPos] != '"'; lex.currPos++ {
		str = append(str, lex.input[lex.currPos])
	}
	if lex.currPos >= len(lex.input) {
		panic(fmt.Sprintf("Probably missing a double quote"))
	}
	return string(str)
}

func (lex *Lexer) readIdentifier() string {
	str := make([]byte, 0)
	for ; lex.currPos < len(lex.input) && isIdentifier(lex.input[lex.currPos]); lex.currPos++ {
		str = append(str, lex.input[lex.currPos])
	}
	lex.currPos--
	return string(str)
}

func isIdentifier(b byte) bool {
	return isDigit(b) || isLetter(b) || b == '_'
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
