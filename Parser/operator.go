package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <operador> → <identificador> | <numeros>

type Operator struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseOperator() (*Operator, error) {
	operator := &Operator{}

	fmt.Println("ParseOperator")
	// either ID or NUM
	tok, lit := p.scanIgnoreWhitespace()
	if tok == Lexer.ID || Lexer.IsNum(tok) {
		operator.AbstractSyntaxTree = append(operator.AbstractSyntaxTree, lit)
	} else {
		return nil, fmt.Errorf("expected identifier or number, got %s", lit)
	}

	return operator, nil

}
