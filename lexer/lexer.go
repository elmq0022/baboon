package lexer

import "github.com/elmq0022/baboon/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
	}

	l.readChar()
	return &l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQUAL, Literal: "=="}
			l.readChar()
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: string(l.ch)}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(l.ch)}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: string(l.ch)}
	case '<':
		tok = token.Token{Type: token.LT, Literal: string(l.ch)}
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOT_EQUAL, Literal: "!="}
			l.readChar()
		} else {
			tok = token.Token{Type: token.NOT, Literal: string(l.ch)}
		}
	case '>':
		tok = token.Token{Type: token.GT, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}

	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '[':
		tok = token.Token{Type: token.LBRACKET, Literal: string(l.ch)}
	case ']':
		tok = token.Token{Type: token.RBRACKET, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}

	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readInt()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.Token{
				Type:    token.ILLEGAL,
				Literal: string(l.ch),
			}
		}
	}
	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) consumeWhitespace() {
	for l.isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
