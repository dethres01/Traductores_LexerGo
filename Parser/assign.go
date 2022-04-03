package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<asignar> → <identificador> = <expresion_arit>

type Assign struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseAssign() (*Assign, error) {
	assign := &Assign{}

	fmt.Println("ParseAssign")
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok != Lexer.ID {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	assign.AbstractSyntaxTree = append(assign.AbstractSyntaxTree, lit)

	tok, lit = p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok != Lexer.ASSIGN {
		return nil, fmt.Errorf("expected =, got %s", lit)
	}
	assign.AbstractSyntaxTree = append(assign.AbstractSyntaxTree, lit)

	exp, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}
	assign.AbstractSyntaxTree = append(assign.AbstractSyntaxTree, exp)

	return assign, nil
}
