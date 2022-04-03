package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<expresion_arit> → (<expresion_arit><operador_arit><expresion_arit>) <rest_arit>
//| <identificador> <rest_arit>
//| <numeros><rest_arit>

type Expression struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseExpression() (*Expression, error) {
	expression := &Expression{}
	fmt.Println("ParseExpression")
	// can either start with ID or NUM or (
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok == Lexer.ID {
		p.unscan()
		fmt.Println("ParseExpression ID")
		identifier, err := p.ParseIdentifier()
		fmt.Println("tok: ", tok, lit)
		if err != nil {
			return nil, err
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, identifier)

	} else if Lexer.IsNum(tok) {
		p.unscan()
		fmt.Println("Number expression")
		tok, lit := p.scanIgnoreWhitespace()
		fmt.Println("tok: ", tok, lit)
		if !Lexer.IsNum(tok) {
			return nil, fmt.Errorf("expected number, got %s", lit)
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, lit)
	} else if tok == Lexer.LPAREN {
		p.unscan()
		// left paren
		fmt.Println("ParseExpression LPAREN")
		tok, lit := p.scanIgnoreWhitespace()
		if tok != Lexer.LPAREN {
			return nil, fmt.Errorf("expected (, got %s", lit)
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, lit)

		// expression
		exp, err := p.ParseExpression()
		if err != nil {
			return nil, err
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, exp)

		// operator
		op, err := p.ParseOperatorArit()
		if err != nil {
			return nil, err
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, op)

		// expression
		exp, err = p.ParseExpression()
		if err != nil {
			return nil, err
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, exp)

		// right paren
		tok, lit = p.scanIgnoreWhitespace()
		if tok != Lexer.RPAREN {
			return nil, fmt.Errorf("expected ), got %s", lit)
		}
		expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, lit)
	} else {
		return nil, fmt.Errorf("expected identifier, number or (, got %s", lit)
	}

	// rest
	fmt.Println("buffer tok: ", p.buf.tok, p.buf.lit)
	rest, err := p.ParseRestExp()
	if err != nil {
		return nil, err
	}
	expression.AbstractSyntaxTree = append(expression.AbstractSyntaxTree, rest)
	return expression, nil
}
