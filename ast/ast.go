package ast

import (
	"github.com/qq52184962/MonkeyCompiler/token"
)

type Program struct {
	Statements []Statement
}

func (p *Program) currentToken() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].currentToken()
	}
	return ""
}

type Node interface {
	currentToken() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) ExpressionNode() {

}

func (i *Identifier) currentToken() string {

}

type LetStatement struct {
	Token token.Token
	Name  Identifier
	Value Expression
}

func (l *LetStatement) ExpressionNode() {

}

func (l *LetStatement) currentToken() string {

}
