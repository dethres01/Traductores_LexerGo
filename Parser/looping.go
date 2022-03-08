package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

func (p *Parser) ParseWhile() (*Statement, error) {
	stmt := &Statement{}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.DO {
		return nil, fmt.Errorf("expected do, got %s", lit)
	}
	stmt.fields = append(stmt.fields, lit)
	block, err := p.Block()
	if err != nil {
		return nil, err
	}
	stmt.fields = append(stmt.fields, block...)
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.WHILE {
		return nil, fmt.Errorf("expected while, got %s", lit)
	}
	stmt.fields = append(stmt.fields, lit)
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.LPAREN {
		return nil, fmt.Errorf("expected (, got %s", lit)
	}
	stmt.fields = append(stmt.fields, lit)
	c, err := p.ParseCondition()
	if err != nil {
		return nil, err
	}
	stmt.condtion = c
	stmt.fields = append(stmt.fields, c.part1, c.cond, c.part2)
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.RPAREN {
		return nil, fmt.Errorf("expected ), got %s", lit)
	}
	stmt.fields = append(stmt.fields, lit)
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.END {
		return nil, fmt.Errorf("expected end, got %s", lit)
	}
	stmt.fields = append(stmt.fields, lit)
	return stmt, nil
}
