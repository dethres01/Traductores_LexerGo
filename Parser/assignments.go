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
	if tok != Lexer.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	stmt := &Assignment{Identifier: lit}
	// now we expect "="
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.ASSIGN {
		return nil, fmt.Errorf("expected =, got %s", lit)
	}
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
		// now we expect a infix operator
		tok, lit = p.scanIgnoreWhitespace()
		if !Lexer.IsInfix(tok) {
			return nil, fmt.Errorf("expected infix operator, got %s", lit)
		}
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
			return stmt, nil
		}
	}
}
