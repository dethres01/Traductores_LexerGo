package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<rest_ordenes> → <orden>; <rest_ordenes> | epsilon

type RestStatements struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseRestStatements() (*RestStatements, error) {
	restStatements := &RestStatements{}
	// we check for a statement, otherwise we return epsilon
	fmt.Println("ParseRestStatements")
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	// probably should do a function to avoid a false positive
	if !Lexer.IsStatement(tok) || tok == Lexer.ELSE {
		// false positive because of or negative operations
		fmt.Println("returning epsilon")
		p.unscan()
	} else {
		fmt.Println("returning statement")
		p.unscan()
		statement, err := p.ParseStatement()
		if err != nil {
			return nil, err
		}
		restStatements.AbstractSyntaxTree = append(restStatements.AbstractSyntaxTree, statement)

		// check for ;
		fmt.Println("ParseRestStatements check for ;")
		tok, lit := p.scanIgnoreWhitespace()
		if tok != Lexer.SEMICOLON {
			return nil, fmt.Errorf("expected ;, got %s", lit)
		}
		restStatements.AbstractSyntaxTree = append(restStatements.AbstractSyntaxTree, lit)

		// check for <rest_ordenes>
		// checks for recursion in the future

		restStatements, err = p.ParseRestStatements()
		if err != nil {
			return nil, err
		}
		restStatements.AbstractSyntaxTree = append(restStatements.AbstractSyntaxTree, restStatements)
	}
	return restStatements, nil
}
