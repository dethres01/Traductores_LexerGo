package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

func (p *Parser) Block() ([]string, error) {
	// LBRACE
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.LBRACE {
		return nil, fmt.Errorf("expected {, got %s", lit)
	}
	// statements
	block := []string{lit}
	for {
		// we could parse different types of statements
		// atm let's just parse assignments
		stmt, err := p.ParseAssignment()
		if err != nil {
			return nil, err
		}
		fmt.Println("got to block again")
		block = append(block, stmt.Fields...)
		tok, lit = p.scanIgnoreWhitespace()
		if tok == Lexer.RBRACE {
			break
		} else {
			p.unscan()
		}
	}
	// RBRACE
	if tok != Lexer.RBRACE {
		return nil, fmt.Errorf("expected }, got %s", lit)
	}
	block = append(block, lit)
	return block, nil
}
