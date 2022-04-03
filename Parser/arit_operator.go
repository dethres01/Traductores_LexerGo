package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<operador_arit> → + | - | * | /

type OperatorArit struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseOperatorArit() (*OperatorArit, error) {
	operatorArit := &OperatorArit{}
	fmt.Println("ParseOperatorArit")
	// either + or - or * or /
	tok, lit := p.scanIgnoreWhitespace()
	if Lexer.IsInfix(tok) {
		operatorArit.AbstractSyntaxTree = append(operatorArit.AbstractSyntaxTree, lit)
	} else {
		return nil, fmt.Errorf("expected OPERATOR, got %s", lit)
	}

	return operatorArit, nil

}
