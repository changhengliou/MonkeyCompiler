package parser

import (
	"fmt"
	"github.com/qq52184962/MonkeyCompiler/ast"
	"github.com/qq52184962/MonkeyCompiler/lexer"
	"github.com/qq52184962/MonkeyCompiler/token"
)

const (
	LOWEST      int = iota
	EQUALS       // ==
	LTORGT       // < or >
	PLUSORMINUS  // + or -
	MULTORDIV    // * or /
	PREFIX       // ! or ++
	FUNCCALL     // function call
)

type Parser struct {
	lex            *lexer.Lexer
	currPos        int
	prefixParseFNs map[string]func() ast.Expression
	infixParseFNs  map[string]func(ast.Expression) ast.Expression
}

func (p *Parser) New() *Parser {
	lex := lexer.New()
	p.prefixParseFNs = make(map[string]func() ast.Expression)
	p.prefixParseFNs[token.IDENTIFIER] = p.parseIdentifier
	p.prefixParseFNs[token.NUMBER] = p.parseNumber
	p.prefixParseFNs[token.MINUS] = p.parsePrefixExpression
	p.prefixParseFNs[token.NOT] = p.parsePrefixExpression
	return &Parser{lex: lex, currPos: 0}
}

func (p *Parser) nextToken() token.Token {
	return p.lex.NextToken()
}

func (p *Parser) peekToken() token.Token {
	return p.lex.PeekToken()
}

func (p *Parser) expectToken(tokType string) token.Token {
	currToken := p.nextToken()
	if currToken.Type != tokType {
		panic(fmt.Sprintf("Expect to be %s, but get %s instead", tokType, currToken.Type))
	}
	return currToken
}

func (p *Parser) Parse() ast.Statement {
	switch p.nextToken().Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseReturnStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{}

	tok := p.expectToken(token.IDENTIFIER)

	statement.Name = &ast.Identifier{Token: tok, Value: tok.Data.(string)}

	p.expectToken(token.ASSIGN)

	p.parseExpression(LOWEST)

	p.expectToken(token.SEMICOLON)

	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{}
	p.parseExpression(LOWEST)
	p.expectToken(token.SEMICOLON)
	return statement
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{}
	p.parseExpression(LOWEST)
	p.expectToken(token.SEMICOLON)
	return statement
}

func (p *Parser) parseIdentifier() ast.Expression {
	tok := p.peekToken()
	return &ast.Identifier{Token: tok, Value: tok.Data.(string)}
}

func (p *Parser) parseNumber() ast.Expression {
	tok := p.peekToken()
	return &ast.NumberLiteral{Token: tok, Value: tok.Data.(float64)}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	tok := p.peekToken()
	expr := ast.InfixExpression{
		Token:    tok,
		Operator: tok.Data.(string),
		Left:     left,
	}
	tok = p.nextToken()
	// get Precedence and replace the LOWEST
	expr.Right = p.parseExpression(LOWEST)
	return expr
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	tok := p.peekToken()
	expr := ast.PrefixExpression{
		Token:    tok,
		Operator: tok.Data.(string),
	}
	tok = p.nextToken()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFNs[p.peekToken().Type]
	if prefix == nil {
		panic("No prefix error")
	}

	leftExpr := prefix()
	// replace this with current token precedence
	precedence := LOWEST
	for p.peekToken().Type != token.SEMICOLON && precedence > LOWEST {
		infix := p.infixParseFNs[p.peekToken().Type]
		if infix == nil {
			return leftExpr
		}

		p.nextToken()
		leftExpr = infix(leftExpr)
	}
	return leftExpr
}
