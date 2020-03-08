package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p* Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type ReturnStatement struct {
	Token token.Token 	// Token.Return
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode() {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }

type LetStatement struct {
	Token token.Token 	// Token.Let
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token 	// Token.Identifier
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
