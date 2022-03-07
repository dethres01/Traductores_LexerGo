package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

type Declaration struct {
	type_of_value string
	identifier    string
}

// this is the function to parse a variable declaration
func (p *Parser) ParseDeclaration() (*Declaration, error) {
	stmt := &Declaration{}
	// our declarations are of the form:
	// var identifier [type] at the moment

	// look for the keyword var
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.VAR {
		return nil, fmt.Errorf("expected var, got %s", lit)
	}
	// atm we are implementing only one identifier
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	stmt.identifier = lit
	// now we need to check if there is a type
	tok, lit = p.scanIgnoreWhitespace()
	if !(tok == Lexer.INT || tok == Lexer.BOOL || tok == Lexer.STRING) {
		return nil, fmt.Errorf("expected type, got %s", lit)
	}
	stmt.type_of_value = lit
	return stmt, nil
}
