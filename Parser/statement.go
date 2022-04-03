package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<orden> → <condicion> | <bucle_while> | <asignar>

type Statement struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseStatement() (*Statement, error) {
	statement := &Statement{}

	fmt.Println("ParseStatement")
	// condition or while loop or assign
	// we have to find out which one it is
	tok, lit := p.scanIgnoreWhitespace()
	if tok == Lexer.IF {
		p.unscan()
		condition, err := p.ParseCondition()
		if err != nil {
			return nil, err
		}
		statement.AbstractSyntaxTree = append(statement.AbstractSyntaxTree, condition)
	} else if tok == Lexer.WHILE {
		p.unscan()
		whileLoop, err := p.ParseWhileLoop()
		if err != nil {
			return nil, err
		}
		statement.AbstractSyntaxTree = append(statement.AbstractSyntaxTree, whileLoop)

	} else if tok == Lexer.ID {
		p.unscan()
		assign, err := p.ParseAssign()
		if err != nil {
			return nil, err
		}
		statement.AbstractSyntaxTree = append(statement.AbstractSyntaxTree, assign)

	} else {
		return nil, fmt.Errorf("expected valid statement, got %s", lit)
	}

	return statement, nil
}
