package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <condicion> → if (<comparación>) <ordenes> <rest_condicion>

func (p *Parser) ParseCondition() (*ASTNode, string, error) {
	condition := &ASTNode{TokenType: Lexer.CONDITION}
	// check for if
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.IF {
		return nil, "", fmt.Errorf("expected if, got %s", lit)
	}
	condition.Children = append(condition.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for (
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.LPAREN {
		return nil, "", fmt.Errorf("expected (, got %s", lit)
	}
	condition.Children = append(condition.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for <comparación>
	comparison, compv, err := p.ParseComparison()
	if err != nil {
		return nil, "", err
	}
	condition.Children = append(condition.Children, *comparison)

	// check for )
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.RPAREN {
		return nil, "", fmt.Errorf("expected ), got %s", lit)
	}
	condition.Children = append(condition.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for <ordenes>
	statements, states, err := p.ParseStatements()
	if err != nil {
		return nil, "", err
	}
	condition.Children = append(condition.Children, *statements)

	// check for <rest_condicion>
	restCondition, rest, err := p.ParseRestCondition()
	if err != nil {
		return nil, "", err
	}
	condition.Children = append(condition.Children, *restCondition)
	result := fmt.Sprintf("if(%s)%s%s", compv, states, rest)
	condition.TokenValue = result

	return condition, condition.TokenValue, nil
}
