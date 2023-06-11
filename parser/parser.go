package parser

import (
    "github.com/gray-adeyi/monkey/ast"
    "github.com/gray-adeyi/monkey/token"
    "github.com/gray-adeyi/monkey/lexer"
)

type Parser struct {
    l *lexer.Lexer

    currToken token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser{
    p := &Parser{l: l}

    // Read two tokens, so currToken and peekToken are both set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.currToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
    program := ast.Program{}
    program.Statements = []ast.Statement{}

    for p.currToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program
}
