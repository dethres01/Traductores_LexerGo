package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<rest_declaraciones> → <declaracion>;<rest_declaracion> | epsilon

type RestDeclarations struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseRestDeclarations() (*RestDeclarations, error) {
	restDeclarations := &RestDeclarations{}

	// this could be either blank(epsilon) or <declaracion>;
	// for performance reasons we could check for type since we know it's a declaration
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("ParseRestDeclarations", tok, lit)
	if !Lexer.IsNum(tok) {
		fmt.Println("ParseRestDeclarations Entered type check")
		// if it's not <declaracion>; then we put the token back
		p.unscan()
	} else {
		// we unscan anyways since we know it's a declaration
		p.unscan()

		// check for <declaracion>
		declaration, err := p.ParseDeclaration()
		if err != nil {
			return nil, err
		}
		restDeclarations.AbstractSyntaxTree = append(restDeclarations.AbstractSyntaxTree, declaration)

		// check for ;
		tok, lit := p.scanIgnoreWhitespace()
		if tok != Lexer.SEMICOLON {
			return nil, fmt.Errorf("expected ;, got %s", lit)
		}
		restDeclarations.AbstractSyntaxTree = append(restDeclarations.AbstractSyntaxTree, lit)

		// check for <rest_declaracion>
		// probably will end up having to go back here since it's recursive
		restDeclaration, err := p.ParseRestDeclarations()
		if err != nil {
			return nil, err
		}
		restDeclarations.AbstractSyntaxTree = append(restDeclarations.AbstractSyntaxTree, restDeclaration)
	}

	return restDeclarations, nil
}
