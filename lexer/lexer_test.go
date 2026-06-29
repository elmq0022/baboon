package lexer

import (
	"testing"

	"github.com/elmq0022/baboon/token"
	"github.com/stretchr/testify/assert"
)

func TestLexerNextToken(t *testing.T) {
	tests := []struct {
		name string
		code string
		want []token.Token
	}{
		{
			name: "assignment",
			code: `=`,
			want: []token.Token{
				{Type: token.ASSIGN, Literal: "="},
			},
		},
		{
			name: "asterisk",
			code: `*`,
			want: []token.Token{
				{Type: token.ASTERISK, Literal: "*"},
			},
		},
		{
			name: "plus",
			code: `+`,
			want: []token.Token{
				{Type: token.PLUS, Literal: "+"},
			},
		},
		{
			name: "minus",
			code: `-`,
			want: []token.Token{
				{Type: token.MINUS, Literal: "-"},
			},
		},
		{
			name: "slash",
			code: `/`,
			want: []token.Token{
				{Type: token.SLASH, Literal: "/"},
			},
		},
		{
			name: "comma",
			code: `,`,
			want: []token.Token{
				{Type: token.COMMA, Literal: ","},
			},
		},
		{
			name: "semicolon",
			code: `;`,
			want: []token.Token{
				{Type: token.SEMICOLON, Literal: ";"},
			},
		},

		{
			name: "parenthesis",
			code: `)()(`,
			want: []token.Token{
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LPAREN, Literal: "("},
			},
		},
		{
			name: "braces",
			code: `}{}{`,
			want: []token.Token{
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.LBRACE, Literal: "{"},
			},
		},
		{
			name: "brackets",
			code: `][][`,
			want: []token.Token{
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.LBRACKET, Literal: "["},
			},
		},

		// keywords
		{
			name: "false",
			code: `false`,
			want: []token.Token{
				{Type: token.FALSE, Literal: "false"},
			},
		},
		{
			name: "function",
			code: `fn`,
			want: []token.Token{
				{Type: token.FUNCTION, Literal: "fn"},
			},
		},
		{
			name: "let",
			code: `let`,
			want: []token.Token{
				{Type: token.LET, Literal: "let"},
			},
		},
		{
			name: "return",
			code: `return`,
			want: []token.Token{
				{Type: token.RETURN, Literal: "return"},
			},
		},
		{
			name: "true",
			code: `true`,
			want: []token.Token{
				{Type: token.TRUE, Literal: "true"},
			},
		},
		{
			name: "int",
			code: `5`,
			want: []token.Token{
				{Type: token.INT, Literal: "5"},
			},
		},
		{
			name: "simple add",
			code: `5 + 5;`,
			want: []token.Token{
				{Type: token.INT, Literal: "5"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
			},
		},
		{
			name: "illegal",
			code: `@@`,
			want: []token.Token{
				{Type: token.ILLEGAL, Literal: "@"},
				{Type: token.ILLEGAL, Literal: "@"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.code)
			got := []token.Token{}
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				got = append(got, tok)
			}
			assert.Equal(t, tt.want, got, "want: %q, got: %q")
		})
	}
}
