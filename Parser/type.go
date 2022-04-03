package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<tipo> → entero | real

type Type struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseType() (*Type, error) {
	type_ := &Type{}
	fmt.Println("ParseType")
	// they are both terminals we check for them
	tok, _ := p.scanIgnoreWhitespace()
	if tok != Lexer.INT && tok != Lexer.FLOAT {
		return nil, fmt.Errorf("expected type (int or float), got %s", tok)
	}
	if tok == Lexer.INT {
		type_.AbstractSyntaxTree = append(type_.AbstractSyntaxTree, "int")
	} else {
		type_.AbstractSyntaxTree = append(type_.AbstractSyntaxTree, "float")
	}

	return type_, nil
}
