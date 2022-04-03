package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <rest_condicion> → end | else <ordenes> end
type restCondition struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseRestCondition() (*restCondition, error) {
	restCondition := &restCondition{}
	fmt.Println("ParseRestCondition")
	// it has to be either end or else
	tok, lit := p.scanIgnoreWhitespace()
	if tok == Lexer.END {
		restCondition.AbstractSyntaxTree = append(restCondition.AbstractSyntaxTree, lit)
	} else if tok == Lexer.ELSE {
		restCondition.AbstractSyntaxTree = append(restCondition.AbstractSyntaxTree, lit)
		// check for <ordenes>
		statements, err := p.ParseStatements()
		if err != nil {
			return nil, err
		}
		restCondition.AbstractSyntaxTree = append(restCondition.AbstractSyntaxTree, statements)

		// check for end
		tok, lit = p.scanIgnoreWhitespace()
		if tok != Lexer.END {
			return nil, fmt.Errorf("expected end, got %s", lit)
		}
		restCondition.AbstractSyntaxTree = append(restCondition.AbstractSyntaxTree, lit)

	} else {
		return nil, fmt.Errorf("expected end or else, got %s", lit)
	}

	return restCondition, nil
}
