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
			name: "=",
			code: `=`,
			want: []token.Token{
				{Type: token.ASSIGN, Literal: "="},
			},
		},
		{
			name: "*",
			code: `*`,
			want: []token.Token{
				{Type: token.ASTERISK, Literal: "*"},
			},
		},
		{
			name: "+",
			code: `+`,
			want: []token.Token{
				{Type: token.PLUS, Literal: "+"},
			},
		},
		{
			name: "-",
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
