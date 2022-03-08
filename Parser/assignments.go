package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

type Assignment struct {
	Identifier string
	Fields     []string
}

func (p *Parser) ParseAssignment() (*Assignment, error) {
	// we are expecting an identifier
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println(tok, lit)

	if tok != Lexer.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	stmt := &Assignment{Identifier: lit}
	stmt.Fields = append(stmt.Fields, lit)
	// now we expect "="
	tok, lit = p.scanIgnoreWhitespace()
	fmt.Println(tok, lit)

	if tok != Lexer.ASSIGN {
		return nil, fmt.Errorf("expected =, got %s", lit)
	}
	stmt.Fields = append(stmt.Fields, lit)
	// now we expect a list of fields
	for {
		// we expect a field, either an IDENT or an INT
		tok, lit = p.scanIgnoreWhitespace()
		if tok == Lexer.EOF {
			return nil, fmt.Errorf("expected field, got EOF")
		}
		if tok == Lexer.IDENT {
			stmt.Fields = append(stmt.Fields, lit)
		} else if tok == Lexer.INT {
			stmt.Fields = append(stmt.Fields, lit)
		} else {
			return nil, fmt.Errorf("expected field, got %s", lit)
		}
		tok, _ = p.scanIgnoreWhitespace()
		if tok == Lexer.EOF {
			return stmt, nil
		} else {
			p.unscan()
		}
		// now we expect a infix operator
		tok, lit = p.scanIgnoreWhitespace()
		if !Lexer.IsInfix(tok) {
			return nil, fmt.Errorf("expected infix operator, got %s", lit)
		}
		stmt.Fields = append(stmt.Fields, lit)
		// now we expect a field, either an IDENT or an INT
		tok, lit = p.scanIgnoreWhitespace()
		if tok == Lexer.EOF {
			return nil, fmt.Errorf("expected field, got EOF")
		}
		if tok == Lexer.IDENT {
			stmt.Fields = append(stmt.Fields, lit)
		} else if tok == Lexer.INT {
			stmt.Fields = append(stmt.Fields, lit)
		} else {
			return nil, fmt.Errorf("expected field, got %s", lit)
		}
		// if eof, we are done
		tok, _ = p.scanIgnoreWhitespace()
		if tok == Lexer.EOF {
			fmt.Println(stmt)
			return stmt, nil
		}
	}
}

func (p *Parser) ParseBool() (*Assignment, error) {
	// look for identifier
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	stmt := &Assignment{}
	stmt.Identifier = lit
	// now we expect "="
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.ASSIGN {
		return nil, fmt.Errorf("expected =, got %s", lit)
	}
	// now we expect a either true or false
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.TRUE && tok != Lexer.FALSE {
		return nil, fmt.Errorf("expected true or false, got %s", lit)
	}
	stmt.Fields = append(stmt.Fields, lit)
	return stmt, nil
}
