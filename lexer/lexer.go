package lexer

import "github.com/kriticalflare/go-brainfuck-yourself/token"

type Lexer struct {
	Tokens []token.Token
	source string
}

func New(source string) Lexer {

	l := Lexer{
		Tokens: []token.Token{},
		source: source,
	}
	l.lex()
	return l
}

func (l *Lexer) lex() {
	for idx, char := range l.source {
		switch char {
		case token.PLUS:
			{
				l.Tokens = append(l.Tokens, token.NewAdd(idx))
			}
		case token.MINUS:
			{
				l.Tokens = append(l.Tokens, token.NewSubtract(idx))
			}
		case token.GT:
			{
				l.Tokens = append(l.Tokens, token.NewNext(idx))
			}
		case token.LT:
			{
				l.Tokens = append(l.Tokens, token.NewPrev(idx))
			}
		case token.LBRACKET:
			{
				l.Tokens = append(l.Tokens, token.NewJumpZero(idx))
			}
		case token.RBRACKET:
			{
				l.Tokens = append(l.Tokens, token.NewJumpNonZero(idx))
			}
		case token.DOT:
			{
				l.Tokens = append(l.Tokens, token.NewOutput(idx))
			}
		case token.COMMA:
			{
				l.Tokens = append(l.Tokens, token.NewInput(idx))
			}
		}
	}
}
