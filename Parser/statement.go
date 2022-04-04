package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<orden> → <condicion> | <bucle_while> | <asignar>

func (p *Parser) ParseStatement() (*ASTNode, string, error) {
	statement := &ASTNode{TokenType: Lexer.STATEMENT}

	// condition or while loop or assign
	// we have to find out which one it is
	tok, lit := p.scanIgnoreWhitespace()
	if tok == Lexer.IF {
		p.unscan()
		condition, condition_v, err := p.ParseCondition()
		if err != nil {
			return nil, "", err
		}
		statement.Children = append(statement.Children, *condition)
		statement.TokenValue = condition_v
	} else if tok == Lexer.WHILE {
		p.unscan()
		whileLoop, whilev, err := p.ParseWhileLoop()
		if err != nil {
			return nil, "", err
		}
		statement.Children = append(statement.Children, *whileLoop)
		statement.TokenValue = whilev

	} else if tok == Lexer.ID {
		p.unscan()
		assign, assignv, err := p.ParseAssign()
		if err != nil {
			return nil, "", err
		}
		statement.Children = append(statement.Children, *assign)
		statement.TokenValue = assignv

	} else {
		return nil, "", fmt.Errorf("expected valid statement, got %s", lit)
	}

	return statement, statement.TokenValue, nil
}
