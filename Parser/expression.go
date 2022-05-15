package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<expresion_arit> → (<expresion_arit><operador_arit><expresion_arit>) <rest_arit>
//| <identificador> <rest_arit>
//| <numeros><rest_arit>

func (p *Parser) ParseExpression() (*ASTNode, string, error) {
	expression := &ASTNode{TokenType: Lexer.EXPRESSION}

	// can either start with ID or NUM or (
	tok, lit := p.scanIgnoreWhitespace()

	if tok == Lexer.ID {
		p.unscan()

		identifier, value, err := p.ParseIdentifier()

		if err != nil {
			return nil, "", err
		}
		expression.Children = append(expression.Children, *identifier)
		expression.TokenValue = value

	} else if Lexer.IsNum(tok) {
		p.unscan()
		tok, lit := p.scanIgnoreWhitespace()
		if !Lexer.IsNum(tok) {
			return nil, "", fmt.Errorf("expected number, got %s", lit)
		}
		expression.Children = append(expression.Children, ASTNode{TokenType: tok, TokenValue: lit})
		value := lit
		expression.TokenValue = value
	} else if tok == Lexer.LPAREN {
		p.unscan()
		// left paren
		tok, lit := p.scanIgnoreWhitespace()
		if tok != Lexer.LPAREN {
			return nil, "", fmt.Errorf("expected (, got %s", lit)
		}
		expression.Children = append(expression.Children, ASTNode{TokenType: tok, TokenValue: lit})

		// expression
		exp, exp_value, err := p.ParseExpression()
		if err != nil {
			return nil, "", err
		}
		expression.Children = append(expression.Children, *exp)
		tok, lit = p.scanIgnoreWhitespace()
		if tok != Lexer.RPAREN {
			return nil, "", fmt.Errorf("expected ), got %s", lit)
		}
		expression.Children = append(expression.Children, ASTNode{TokenType: tok, TokenValue: lit})
		value := fmt.Sprintf("(%s)", exp_value)
		expression.TokenValue = value
	} else {
		return nil, "", fmt.Errorf("expected identifier, number or (, got %s", lit)
	}

	// rest
	rest, restexp_value, err := p.ParseRestExp()
	if err != nil {
		return nil, "", err
	}
	expression.Children = append(expression.Children, *rest)
	result := fmt.Sprintf("%s %s", expression.TokenValue, restexp_value)
	expression.TokenValue = result
	return expression, expression.TokenValue, nil
}
