package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

type Condition struct {
	part1 string
	part2 string
	cond  string
}

func (p *Parser) ParseCondition() (*Condition, error) {
	cond := &Condition{}
	tok, lit := p.scanIgnoreWhitespace()
	if !(tok == Lexer.IDENT || tok == Lexer.INT) {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	cond.part1 = lit
	tok, lit = p.scanIgnoreWhitespace()
	fmt.Println(tok, lit)
	if !Lexer.IsComparative(tok) {
		return nil, fmt.Errorf("expected conditional, got %s", lit)
	}
	cond.cond = lit
	tok, lit = p.scanIgnoreWhitespace()
	if !(tok == Lexer.IDENT || tok == Lexer.INT) {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	cond.part2 = lit
	return cond, nil
}
