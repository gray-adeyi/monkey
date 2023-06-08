package token

const (
    ILLEGAL = "ILLEGAL"
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

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}
