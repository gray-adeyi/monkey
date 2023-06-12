package ast

import (
    "github.com/gray-adeyi/monkey/token"
)

type Node interface {
    TokenLiteral() string
}

type Statement interface{
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

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

type LetStatement struct {
    Token token.Token  // The token token.LET
    Name *Identifier
    Value Expression
}

func (ls *LetStatement)statementNode() {}
func (ls *LetStatement)TokenLiteral() string{ return ls.Token.Literal}


type Identifier struct{
    Token token.Token // The token token.IDENT
    Value string
}

func (i *Identifier) expressionNode(){}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct{
    Token token.Token // the 'return' token
    ReturnValue Expression
}

func (rs *ReturnStatement) statementNode(){}
func (rs *ReturnStatement) TokenLiteral() string{
    return rs.Token.Literal
}
