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
    FUNCTION = "FUNCTION"
    LET = "LET"
)

var keywords =  map[string]TokenType{
    "fn" : FUNCTION,
    "let" : LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}



type Token struct {
    Type TokenType
    Literal string
}
