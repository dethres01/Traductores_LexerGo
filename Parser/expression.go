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
	fmt.Println("ParseExpression")
	// can either start with ID or NUM or (
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println("tok: ", tok, lit)
	if tok == Lexer.ID {
		p.unscan()
		fmt.Println("ParseExpression ID")
		identifier, value, err := p.ParseIdentifier()
		fmt.Println("tok: ", tok, lit)
		if err != nil {
			return nil, "", err
		}
		expression.Children = append(expression.Children, *identifier)
		expression.TokenValue = value

	} else if Lexer.IsNum(tok) {
		p.unscan()
		fmt.Println("Number expression")
		tok, lit := p.scanIgnoreWhitespace()
		fmt.Println("tok: ", tok, lit)
		if !Lexer.IsNum(tok) {
			return nil, "", fmt.Errorf("expected number, got %s", lit)
		}
		expression.Children = append(expression.Children, ASTNode{TokenType: tok, TokenValue: lit})
		value := lit
		expression.TokenValue = value
	} else if tok == Lexer.LPAREN {
		p.unscan()
		// left paren
		fmt.Println("ParseExpression LPAREN")
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

		// operator
		op, arit_value, err := p.ParseOperatorArit()
		if err != nil {
			return nil, "", err
		}
		expression.Children = append(expression.Children, *op)

		// expression
		exp2, exp2_value, err := p.ParseExpression()
		if err != nil {
			return nil, "", err
		}
		expression.Children = append(expression.Children, *exp2)

		// right paren
		tok, lit = p.scanIgnoreWhitespace()
		if tok != Lexer.RPAREN {
			return nil, "", fmt.Errorf("expected ), got %s", lit)
		}
		expression.Children = append(expression.Children, ASTNode{TokenType: tok, TokenValue: lit})
		value := fmt.Sprintf("(%s %s %s)", exp_value, arit_value, exp2_value)
		expression.TokenValue = value
	} else {
		return nil, "", fmt.Errorf("expected identifier, number or (, got %s", lit)
	}

	// rest
	fmt.Println("buffer tok: ", p.buf.tok, p.buf.lit)
	rest, restexp_value, err := p.ParseRestExp()
	if err != nil {
		return nil, "", err
	}
	expression.Children = append(expression.Children, *rest)
	result := fmt.Sprintf("%s %s", expression.TokenValue, restexp_value)
	expression.TokenValue = result
	return expression, expression.TokenValue, nil
}
