package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <rest_arit> → <operador_arit><expresion_arit><rest_arit> | epsilon

type RestExp struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseRestExp() (*RestExp, error) {
	restExp := &RestExp{}
	fmt.Println("ParseRestExp")
	fmt.Println("buffer tok: ", p.buf.tok, p.buf.lit)
	// blank or infix operator
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if Lexer.IsInfix(tok) {
		p.unscan()
		fmt.Println("ParseRestExp Infix Entered")
		// <operador_arit>
		operatorArit, err := p.ParseOperatorArit()
		if err != nil {
			return nil, err
		}
		restExp.AbstractSyntaxTree = append(restExp.AbstractSyntaxTree, operatorArit)

		// <expresion_arit>
		exp, err := p.ParseExpression()
		if err != nil {
			return nil, err
		}
		restExp.AbstractSyntaxTree = append(restExp.AbstractSyntaxTree, exp)

		// <rest_arit>
		rest, err := p.ParseRestExp()
		if err != nil {
			return nil, err
		}
		restExp.AbstractSyntaxTree = append(restExp.AbstractSyntaxTree, rest)
	} else {
		p.unscan()
		fmt.Println("ParseRestExp Epsilon Entered")
		fmt.Println("buf: ", p.buf.tok, p.buf.lit)
	}

	return restExp, nil
}
