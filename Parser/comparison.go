package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <comparacion> → <operador> <condicion_op> <operador>

func (p *Parser) ParseComparison() (*ASTNode, string, error) {
	comparison := &ASTNode{TokenType: Lexer.COMPARISON}
	fmt.Println("ParseComparison")

	// check for <operador>
	operador, op1, err := p.ParseOperator()
	if err != nil {
		return nil, "", err
	}
	comparison.Children = append(comparison.Children, *operador)

	// check for <condicion_op>
	condicionOp, cond, err := p.ParseConditionOp()
	if err != nil {
		return nil, "", err
	}

	comparison.Children = append(comparison.Children, *condicionOp)

	// check for <operador>
	operador2, op2, err := p.ParseOperator()
	if err != nil {
		return nil, "", err
	}
	comparison.Children = append(comparison.Children, *operador2)
	result := fmt.Sprintf("%s %s %s", op1, cond, op2)
	comparison.TokenValue = result

	return comparison, comparison.TokenValue, nil
}
