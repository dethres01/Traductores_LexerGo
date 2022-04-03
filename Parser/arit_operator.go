package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<operador_arit> → + | - | * | /

func (p *Parser) ParseOperatorArit() (*ASTNode, string, error) {
	operatorArit := &ASTNode{TokenType: Lexer.OPERATOR_ARIT}
	fmt.Println("ParseOperatorArit")
	// either + or - or * or /
	tok, lit := p.scanIgnoreWhitespace()
	if Lexer.IsInfix(tok) {
		operatorArit.TokenValue = lit
		operatorArit.Children = append(operatorArit.Children, ASTNode{TokenType: tok, TokenValue: lit})
	} else {
		return nil, "", fmt.Errorf("expected OPERATOR, got %s", lit)
	}

	return operatorArit, operatorArit.TokenValue, nil

}
