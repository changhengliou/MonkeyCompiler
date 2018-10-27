package parser

import (
	"github.com/qq52184962/MonkeyCompiler/ast"
	"github.com/qq52184962/MonkeyCompiler/lexer"
	"github.com/qq52184962/MonkeyCompiler/token"
)

type Parser struct {
	lex     *lexer.Lexer
	currPos int
}

func (p *Parser) New() *Parser {
	lex := lexer.New()
	return &Parser{lex: lex, currPos: 0}
}

func (p *Parser) nextToken() token.Token {
	return p.lex.NextToken()
}

func Parse() *ast.Program {
	return nil
}
