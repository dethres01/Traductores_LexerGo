package Parser

import "fmt"

// <comparacion> → <operador> <condicion_op> <operador>

type Comparison struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseComparison() (*Comparison, error) {
	comparison := &Comparison{}
	fmt.Println("ParseComparison")

	// check for <operador>
	operador, err := p.ParseOperator()
	if err != nil {
		return nil, err
	}
	comparison.AbstractSyntaxTree = append(comparison.AbstractSyntaxTree, operador)

	// check for <condicion_op>
	condicionOp, err := p.ParseConditionOp()
	if err != nil {
		return nil, err
	}
	comparison.AbstractSyntaxTree = append(comparison.AbstractSyntaxTree, condicionOp)

	// check for <operador>
	operador, err = p.ParseOperator()
	if err != nil {
		return nil, err
	}
	comparison.AbstractSyntaxTree = append(comparison.AbstractSyntaxTree, operador)

	return comparison, nil
}
