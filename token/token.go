package token

type TokenKind rune

const (
	PLUS     = '+'
	MINUS    = '-'
	GT       = '>'
	LT       = '<'
	LBRACKET = '['
	RBRACKET = ']'
	DOT      = '.'
	COMMA    = ','
)

type Token interface {
	Kind() TokenKind
	SourcePosition() uint
}

type Add struct {
	sourcePosition uint
}

func NewAdd(sourcePosition int) *Add {
	return &Add{sourcePosition: uint(sourcePosition)}
}

func (a *Add) Kind() TokenKind {
	return PLUS
}

func (a *Add) SourcePosition() uint {
	return a.sourcePosition
}

type Subtract struct {
	sourcePosition uint
}

func NewSubtract(sourcePosition int) *Subtract {
	return &Subtract{sourcePosition: uint(sourcePosition)}
}

func (s *Subtract) Kind() TokenKind {
	return MINUS
}

func (s *Subtract) SourcePosition() uint {
	return s.sourcePosition
}

type Next struct {
	sourcePosition uint
}

func NewNext(sourcePosition int) *Next {
	return &Next{sourcePosition: uint(sourcePosition)}
}

func (n *Next) Kind() TokenKind {
	return LT
}

func (n *Next) SourcePosition() uint {
	return n.sourcePosition
}

type Prev struct {
	sourcePosition uint
}

func NewPrev(sourcePosition int) *Prev {
	return &Prev{sourcePosition: uint(sourcePosition)}
}

func (p *Prev) Kind() TokenKind {
	return MINUS
}

func (p *Prev) SourcePosition() uint {
	return p.sourcePosition
}

type JumpZero struct {
	Target         uint
	sourcePosition uint
}

func NewJumpZero(sourcePosition int) *JumpZero {
	return &JumpZero{sourcePosition: uint(sourcePosition)}
}

func (jz *JumpZero) Kind() TokenKind {
	return LBRACKET
}

func (jz *JumpZero) SourcePosition() uint {
	return jz.sourcePosition
}

type JumpNonZero struct {
	Target         uint
	sourcePosition uint
}

func NewJumpNonZero(sourcePosition int) *JumpNonZero {
	return &JumpNonZero{sourcePosition: uint(sourcePosition)}
}

func (jnz *JumpNonZero) Kind() TokenKind {
	return RBRACKET
}

func (jnz *JumpNonZero) SourcePosition() uint {
	return jnz.sourcePosition
}

type Output struct {
	sourcePosition uint
}

func NewOutput(sourcePosition int) *Output {
	return &Output{sourcePosition: uint(sourcePosition)}
}

func (o *Output) Kind() TokenKind {
	return DOT
}

func (o *Output) SourcePosition() uint {
	return o.sourcePosition
}

type Input struct {
	sourcePosition uint
}

func NewInput(sourcePosition int) *Input {
	return &Input{sourcePosition: uint(sourcePosition)}
}

func (i *Input) Kind() TokenKind {
	return COMMA
}

func (i *Input) SourcePosition() uint {
	return i.sourcePosition
}
