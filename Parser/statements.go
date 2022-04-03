package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <ordenes> → <orden> ; <rest_ordenes>

type Statements struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseStatements() (*Statements, error) {
	statements := &Statements{}

	// check for <orden>
	fmt.Println("ParseStatements")
	statement, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}
	statements.AbstractSyntaxTree = append(statements.AbstractSyntaxTree, statement)

	// check for ;
	fmt.Println("ParseStatements check for ;")

	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok != Lexer.SEMICOLON {
		return nil, fmt.Errorf("expected ;, got %s", lit)
	}
	statements.AbstractSyntaxTree = append(statements.AbstractSyntaxTree, lit)

	// check for <rest_ordenes>
	restStatements, err := p.ParseRestStatements()
	if err != nil {
		return nil, err
	}
	statements.AbstractSyntaxTree = append(statements.AbstractSyntaxTree, restStatements)

	return statements, nil
}
