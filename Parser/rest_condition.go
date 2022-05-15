package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <rest_condicion> → end | else <ordenes> end

func (p *Parser) ParseRestCondition() (*ASTNode, string, error) {
	restCondition := &ASTNode{TokenType: Lexer.REST_CONDITION}
	// it has to be either end or else
	tok, lit := p.scanIgnoreWhitespace()
	if tok == Lexer.END {
		restCondition.TokenValue = lit
		restCondition.Children = append(restCondition.Children, ASTNode{TokenType: tok, TokenValue: lit})
	} else if tok == Lexer.ELSE {
		// add goto
		p.ic.if_condition(lit)
		else_token := lit
		restCondition.Children = append(restCondition.Children, ASTNode{TokenType: tok, TokenValue: lit})
		// check for <ordenes>
		statements, statements_value, err := p.ParseStatements()
		if err != nil {
			return nil, "", err
		}
		p.ic.EndIf(lit)
		restCondition.Children = append(restCondition.Children, *statements)

		// check for end
		tok, lit = p.scanIgnoreWhitespace()
		if tok != Lexer.END {
			return nil, "", fmt.Errorf("expected end, got %s", lit)
		}
		restCondition.Children = append(restCondition.Children, ASTNode{TokenType: tok, TokenValue: lit})
		result := fmt.Sprintf("%s %s %s", else_token, statements_value, lit)
		restCondition.TokenValue = result

	} else {
		return nil, "", fmt.Errorf("expected end or else, got %s", lit)
	}

	return restCondition, restCondition.TokenValue, nil
}
