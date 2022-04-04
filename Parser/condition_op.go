package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<condicion_op> → == | < | > | <= | >= | <>

func (p *Parser) ParseConditionOp() (*ASTNode, string, error) {
	condition_op := &ASTNode{TokenType: Lexer.CONDITION_OP}
	// it has to be either ==, <, >, <=, >= or <>
	tok, lit := p.scanIgnoreWhitespace()
	if Lexer.IsComparative(tok) {
		condition_op.TokenValue = lit
		condition_op.Children = append(condition_op.Children, ASTNode{TokenType: tok, TokenValue: lit})
	} else {
		return nil, "", fmt.Errorf("expected comparative operator, got %s", lit)
	}

	return condition_op, condition_op.TokenValue, nil

}
