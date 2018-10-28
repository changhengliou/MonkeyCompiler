package token

type Token struct {
	Type     string
	Data     interface{}
	Line     int
	Position int
}

const (
	// special purpose
	EOF     = "EOF"
	EOL     = "\n"
	ILLEGAL = "illegal"

	// keyword
	FUNC   = "func"
	LET    = "let"
	RETURN = "return"
	IF     = "if"
	ELSE   = "else"
	FOR    = "FOR"
	TRUE   = "true"
	FALSE  = "false"

	// identifier
	IDENTIFIER = "identifier"
	LITERAL    = "literal"

	// data type
	STRINGLITERAL = "string"
	NUMBER        = "number"

	// special tokens
	LPAREN      = "("
	RPAREN      = ")"
	LBRACKET    = "{"
	RBRACKET    = "}"
	COMMA       = ","
	SEMICOLON   = ";"
	COLON       = ":"
	LSQRBRACKET = "["
	RSQRBRACKET = "]"

	// operator
	ASSIGN   = "="
	NOT      = "!"
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	GT  = ">"
	LT  = "<"

	EQUAL    = "=="
	NOTEQUAL = "!="
)

var keywords = map[string]string{
	"func":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

func GetKeywordOrIdentifier(str string) string {
	if s, ok := keywords[str]; ok {
		return s
	}
	return IDENTIFIER
}
