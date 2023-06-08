package token

type TokenType string

const (
    ILLEGAL TokenType = "ILLEGAL"
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foo, x, y,...
    INT = "INT" // 12345

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimeters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION
    LET
)



type Token struct {
    Type TokenType
    Literal string
}
