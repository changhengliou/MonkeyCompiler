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
	ASSIGN      = "="
	GREATER     = ">"
	SMALLER     = "<"
	LSQRBRACKET = "["
	RSQRBRACKET = "]"

	// operator
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
)

var keywords = map[string]string{
	"fn":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

func GetIdentifier(str string) string {
	if s, ok := keywords[str]; ok {
		return s
	}
	return ILLEGAL
}
