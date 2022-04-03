package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <condicion> → if (<comparación>) <ordenes> <rest_condicion>

type Condition struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseCondition() (*Condition, error) {
	condition := &Condition{}
	fmt.Println("ParseCondition")
	// check for if
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.IF {
		return nil, fmt.Errorf("expected if, got %s", lit)
	}
	condition.AbstractSyntaxTree = append(condition.AbstractSyntaxTree, lit)

	// check for (
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.LPAREN {
		return nil, fmt.Errorf("expected (, got %s", lit)
	}
	condition.AbstractSyntaxTree = append(condition.AbstractSyntaxTree, lit)

	// check for <comparación>
	comparison, err := p.ParseComparison()
	if err != nil {
		return nil, err
	}
	condition.AbstractSyntaxTree = append(condition.AbstractSyntaxTree, comparison)

	// check for )
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.RPAREN {
		return nil, fmt.Errorf("expected ), got %s", lit)
	}
	condition.AbstractSyntaxTree = append(condition.AbstractSyntaxTree, lit)

	// check for <ordenes>
	statements, err := p.ParseStatements()
	if err != nil {
		return nil, err
	}
	condition.AbstractSyntaxTree = append(condition.AbstractSyntaxTree, statements)

	// check for <rest_condicion>
	restCondition, err := p.ParseRestCondition()
	if err != nil {
		return nil, err
	}
	condition.AbstractSyntaxTree = append(condition.AbstractSyntaxTree, restCondition)

	return condition, nil
}
