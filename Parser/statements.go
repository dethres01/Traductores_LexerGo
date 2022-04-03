package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <ordenes> → <orden> ; <rest_ordenes>

func (p *Parser) ParseStatements() (*ASTNode, string, error) {
	statements := &ASTNode{TokenType: Lexer.STATEMENTS}

	// check for <orden>
	fmt.Println("ParseStatements")
	statement, statement_value, err := p.ParseStatement()
	if err != nil {
		return nil, "", err
	}
	statements.Children = append(statements.Children, *statement)

	// check for ;
	fmt.Println("ParseStatements check for ;")

	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok != Lexer.SEMICOLON {
		return nil, "", fmt.Errorf("expected ;, got %s", lit)
	}
	statements.Children = append(statements.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for <rest_ordenes>
	restStatements, restStatements_value, err := p.ParseRestStatements()
	if err != nil {
		return nil, "", err
	}
	statements.Children = append(statements.Children, *restStatements)
	result := fmt.Sprintf("%s;%s", statement_value, restStatements_value)

	statements.TokenValue = result

	return statements, statements.TokenValue, nil
}
