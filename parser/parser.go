package parser

import (
	"fmt"

	"github.com/kriticalflare/go-brainfuck-yourself/lexer"
	"github.com/kriticalflare/go-brainfuck-yourself/token"
)

type Parser struct {
	Tokens []token.Token
	stack  []stackElement
	Errors []string
}

type stackElement struct {
	token    *token.JumpZero
	position int
}

func New(l lexer.Lexer) Parser {
	p := Parser{
		Tokens: l.Tokens,
		stack:  []stackElement{},
	}
	return p
}

func (p *Parser) CheckParserErrors() bool {
	return len(p.Errors) != 0
}

func (p *Parser) ParseProgram() {
	for idx, tok := range p.Tokens {
		switch t := tok.(type) {
		case *token.JumpZero:
			{
				p.stack = append(p.stack, stackElement{token: t, position: idx})
			}
		case *token.JumpNonZero:
			{
				stackLen := len(p.stack)
				if stackLen == 0 {
					p.Errors = append(p.Errors, fmt.Sprintf("char:%d Matching [ not found\n", t.SourcePosition()))
					continue
				}

				top := p.stack[stackLen-1]
				jz := top.token
				jz.Target = uint(idx)
				t.Target = uint(top.position)

				p.stack = p.stack[:len(p.stack)-1]
			}
		}
	}

	for _, ele := range p.stack {
		p.Errors = append(p.Errors, fmt.Sprintf("char:%d Matching ] not found\n", ele.token.SourcePosition()))
	}
}
