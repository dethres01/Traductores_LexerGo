package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <operador> → <identificador> | <numeros>

func (p *Parser) ParseOperator() (*ASTNode, string, error) {
	operator := &ASTNode{TokenType: Lexer.OPERATOR}

	// either ID or NUM
	tok, lit := p.scanIgnoreWhitespace()
	if tok == Lexer.ID || Lexer.IsNum(tok) {
		operator.TokenValue = lit

		operator.Children = append(operator.Children, ASTNode{TokenType: tok, TokenValue: lit})
	} else {
		return nil, "", fmt.Errorf("expected identifier or number, got %s", lit)
	}

	return operator, operator.TokenValue, nil

}
