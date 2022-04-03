package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<rest_declaraciones> → <declaracion>;<rest_declaracion> | epsilon

func (p *Parser) ParseRestDeclarations() (*ASTNode, string, error) {
	restDeclarations := &ASTNode{TokenType: Lexer.REST_DECLARATIONS}

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
		declaration, declaration_value, err := p.ParseDeclaration()
		if err != nil {
			return nil, "", err
		}
		// add child to the declarations
		restDeclarations.Children = append(restDeclarations.Children, *declaration)

		// check for ;
		tok, lit := p.scanIgnoreWhitespace()
		if tok != Lexer.SEMICOLON {
			return nil, "", fmt.Errorf("expected ;, got %s", lit)
		}
		restDeclarations.Children = append(restDeclarations.Children, ASTNode{TokenType: tok, TokenValue: lit})

		// check for <rest_declaracion>
		// probably will end up having to go back here since it's recursive
		restDeclaration_r, value, err := p.ParseRestDeclarations()
		if err != nil {
			return nil, "", err
		}
		result := fmt.Sprintf("%s;%s", declaration_value, value)
		restDeclarations.Children = append(restDeclarations.Children, *restDeclaration_r)
		restDeclarations.TokenValue = result

	}

	return restDeclarations, restDeclarations.TokenValue, nil
}
