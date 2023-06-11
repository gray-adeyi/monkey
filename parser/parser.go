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
    program := &ast.Program{}
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

func (p *Parser) parseStatement() ast.Statement{
    switch p.currToken.Type {
    case token.LET:
        return p.parseLetStatement()
    default:
        return nil
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement{
    stmt := &ast.LetStatement{Token: p.currToken}

    if !p.expectPeek(token.IDENT){
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

    if !p.expectPeek(token.ASSIGN){
        return nil
    }

    // TODO: We're skipping the expressions until we
    // encounter a semicolon
    for !p.currTokenIs(token.SEMICOLON){
        p.nextToken()
    }
    return stmt
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
    return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser)expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        return false
    }
}
