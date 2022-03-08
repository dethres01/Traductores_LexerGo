package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

type Statement struct {
	condtion *Condition
	fields   []string
}

// if statements are of the form
func (p *Parser) ParseIf() (*Statement, error) {
	// create a new statement
	stmt := &Statement{}
	// we expect an if
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.IF {
		return nil, fmt.Errorf("expected if, got %s", lit)
	}
	// add the if to the statement
	stmt.fields = append(stmt.fields, lit)
	// now we expect a condition
	c, err := p.ParseCondition()
	if err != nil {
		return nil, err
	}
	stmt.condtion = c
	stmt.fields = append(stmt.fields, c.part1, c.cond, c.part2)
	// now we expect a then
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.THEN {
		return nil, fmt.Errorf("expected then, got %s", lit)
	}
	// add the then to the statement
	stmt.fields = append(stmt.fields, lit)
	// now we expect a block
	block, err := p.Block()
	if err != nil {
		return nil, err
	}
	stmt.fields = append(stmt.fields, block...)
	// now we expect an else
	tok, lit = p.scanIgnoreWhitespace()
	if tok == Lexer.ELSE {
		stmt.fields = append(stmt.fields, lit)
		// now we expect a block
		block, err := p.Block()
		if err != nil {
			return nil, err
		}
		fmt.Println("did block on else")
		stmt.fields = append(stmt.fields, block...)
		fmt.Println(stmt.fields)
		tok, lit = p.scanIgnoreWhitespace()
		if tok != Lexer.END {
			return nil, fmt.Errorf("expected end, got %s", lit)
		}
		stmt.fields = append(stmt.fields, lit)
		return stmt, nil

	} else if tok != Lexer.END {
		return nil, fmt.Errorf("expected end, got %s", lit)
	} else {
		stmt.fields = append(stmt.fields, lit)
		return stmt, nil
	}

}
