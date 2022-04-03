package Parser

import (
	"AnalisisLexico/Lexer"
)

type Parser struct {
	s   *Lexer.Scanner
	buf struct {
		tok Lexer.Token
		lit string
		n   int
	}
}

func NewParser(s *Lexer.Scanner) *Parser {
	return &Parser{s: s}
}

func (p *Parser) scan() (tok Lexer.Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}
	tok, lit = p.s.Scan()
	p.buf.tok, p.buf.lit = tok, lit
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }

//helper function to scan without whitespace
// might not need it
func (p *Parser) scanIgnoreWhitespace() (tok Lexer.Token, lit string) {
	tok, lit = p.scan()
	for tok == Lexer.WS {
		tok, lit = p.scan()
	}

	return
}
