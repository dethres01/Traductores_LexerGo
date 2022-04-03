package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

type Declarations struct {
	AbstractSyntaxTree []interface{}
}

//<declaraciones> → <declaracion>**;**<rest_declaracion>
// this is the function to parse a variable declaration
func (p *Parser) ParseDeclarations() (*Declarations, error) {
	declarations := &Declarations{}

	fmt.Println("ParseDeclarations")

	// check for <declaracion>
	declaration, err := p.ParseDeclaration()
	if err != nil {
		return nil, err
	}
	// print the declaration information to the console
	fmt.Println("declaration: ", declaration.AbstractSyntaxTree)
	declarations.AbstractSyntaxTree = append(declarations.AbstractSyntaxTree, declaration)

	// check for **;**
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok != Lexer.SEMICOLON {
		return nil, fmt.Errorf("expected ;, got %s", lit)
	}
	declarations.AbstractSyntaxTree = append(declarations.AbstractSyntaxTree, lit)

	// check for <rest_declaracion>
	restDeclaration, err := p.ParseRestDeclarations()

	if err != nil {
		return nil, err
	}
	declarations.AbstractSyntaxTree = append(declarations.AbstractSyntaxTree, restDeclaration)

	return declarations, nil
}
