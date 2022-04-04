package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<rest_ordenes> → <orden>; <rest_ordenes> | epsilon

func (p *Parser) ParseRestStatements() (*ASTNode, string, error) {
	restStatements := &ASTNode{TokenType: Lexer.REST_STATEMENTS}
	// we check for a statement, otherwise we return epsilon

	tok, _ := p.scanIgnoreWhitespace()

	// probably should do a function to avoid a false positive
	if !Lexer.IsStatement(tok) || tok == Lexer.ELSE {
		// false positive because of or negative operations
		p.unscan()
	} else {
		p.unscan()
		statement, statement_value, err := p.ParseStatement()
		if err != nil {
			return nil, "", err
		}
		restStatements.Children = append(restStatements.Children, *statement)

		// check for ;
		tok, lit := p.scanIgnoreWhitespace()
		if tok != Lexer.SEMICOLON {
			return nil, "", fmt.Errorf("expected ;, got %s", lit)
		}
		restStatements.Children = append(restStatements.Children, ASTNode{TokenType: tok, TokenValue: lit})

		// check for <rest_ordenes>
		// checks for recursion in the future

		restStatements_r, value, err := p.ParseRestStatements()
		if err != nil {
			return nil, "", err
		}
		result := fmt.Sprintf("%s;%s", statement_value, value)
		restStatements.Children = append(restStatements.Children, *restStatements_r)
		restStatements.TokenValue = result
	}
	return restStatements, restStatements.TokenValue, nil
}
