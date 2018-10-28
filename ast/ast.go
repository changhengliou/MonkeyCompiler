package ast

import (
	"github.com/qq52184962/MonkeyCompiler/token"
)

type Program struct {
	Statements []Statement
}

func (p *Program) currentToken() interface{} {
	if len(p.Statements) > 0 {
		return p.Statements[0].currentToken()
	}
	return ""
}

type Node interface {
	currentToken() interface{}
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type NumberLiteral struct {
	Token token.Token
	Value float64
}

func (n NumberLiteral) ExpressionNode() {

}

func (n NumberLiteral) currentToken() interface{} {
	return n.Token.Data
}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (ie InfixExpression) ExpressionNode() {

}

func (ie InfixExpression) currentToken() interface{} {
	return ie.Token.Data
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe PrefixExpression) ExpressionNode() {

}

func (pe PrefixExpression) currentToken() interface{} {
	return pe.Token.Data
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) ExpressionNode() {

}

func (i Identifier) currentToken() interface{} {
	return i.Token.Data
}

type LetStatement struct {
	Token *token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode() {
}

func (ls *LetStatement) currentToken() interface{} {
	return ls.Token.Data
}

type ReturnStatement struct {
	Token *token.Token
	Value Expression
}

func (rs *ReturnStatement) StatementNode() {

}

func (rs *ReturnStatement) currentToken() interface{} {
	return rs.Token.Data
}

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (es *ExpressionStatement) StatementNode() {

}

func (es *ExpressionStatement) currentToken() interface{} {
	return es.Token.Data
}
