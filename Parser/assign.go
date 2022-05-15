package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<asignar> → <identificador> = <expresion_arit>

func (p *Parser) ParseAssign() (*ASTNode, string, error) {
	assign := &ASTNode{TokenType: Lexer.ASSIGNMENT}

	tok, lit := p.scanIgnoreWhitespace()

	if tok != Lexer.ID {
		return nil, "", fmt.Errorf("expected identifier, got %s", lit)
	}
	id_value := lit
	assign.Children = append(assign.Children, ASTNode{TokenType: tok, TokenValue: lit})

	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.ASSIGN {
		return nil, "", fmt.Errorf("expected =, got %s", lit)
	}
	assign.Children = append(assign.Children, ASTNode{TokenType: tok, TokenValue: lit})

	exp, exp_value, err := p.ParseExpression()
	if err != nil {
		return nil, "", err
	}
	p.ic.Assignment(id_value, exp_value)
	assign.Children = append(assign.Children, *exp)

	result := fmt.Sprintf("%s = %s", id_value, exp_value)
	assign.TokenValue = result

	return assign, assign.TokenValue, nil
}
