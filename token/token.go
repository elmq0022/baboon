package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// single char operators
	ASSIGN   = "ASSIGN"
	ASTERISK = "ASTERISK"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	SLASH    = "SLASH"

	// grouping symbols
	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	// values
	INT = "INT"

	// identifiers
	IDENT = "IDENT"

	// language keywords
	FALSE    = "FALSE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
}

func LookupKeyword(s string) TokenType {
	keyword, ok := keywords[s]
	if ok {
		return keyword
	}
	return IDENT
}
