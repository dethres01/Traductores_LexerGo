package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<condicion_op> → == | < | > | <= | >= | <>

type ConditionOp struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseConditionOp() (*ConditionOp, error) {
	condition_op := &ConditionOp{}
	fmt.Println("ParseConditionOp")
	// it has to be either ==, <, >, <=, >= or <>
	tok, lit := p.scanIgnoreWhitespace()
	if Lexer.IsComparative(tok) {
		condition_op.AbstractSyntaxTree = append(condition_op.AbstractSyntaxTree, lit)
	} else {
		return nil, fmt.Errorf("expected comparative operator, got %s", lit)
	}

	return condition_op, nil

}
