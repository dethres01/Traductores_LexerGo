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
func (p *Parser) ParseDeclarations() (*ASTNode, string, error) {
	declarations := &ASTNode{TokenType: Lexer.DECLARATIONS}

	// check for <declaracion>
	declaration, declaration_value, err := p.ParseDeclaration()
	if err != nil {
		return nil, "", err
	}
	// print the declaration information to the console

	// add child to the declarations
	declarations.Children = append(declarations.Children, *declaration)

	// check for **;**
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.SEMICOLON {
		return nil, "", fmt.Errorf("expected ;, got %s", lit)
	}
	// add child to the declarations
	declarations.Children = append(declarations.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for <rest_declaracion>
	restDeclaration, restdeclaration_value, err := p.ParseRestDeclarations()

	if err != nil {
		return nil, "", err
	}
	// add child to the declarations
	declarations.Children = append(declarations.Children, *restDeclaration)

	result := fmt.Sprintf("%s;%s", declaration_value, restdeclaration_value)
	declarations.TokenValue = result

	return declarations, declarations.TokenValue, nil
}
