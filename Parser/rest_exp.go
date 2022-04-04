package Parser

import (
	"AnalisisLexico/Lexer"
)

// <rest_arit> → <operador_arit><expresion_arit><rest_arit> | epsilon

func (p *Parser) ParseRestExp() (*ASTNode, string, error) {
	restExp := &ASTNode{TokenType: Lexer.REST_EXP}

	// blank or infix operator
	tok, _ := p.scanIgnoreWhitespace()

	if Lexer.IsInfix(tok) {
		p.unscan()
		// <operador_arit>
		operatorArit, operatorArit_value, err := p.ParseOperatorArit()
		if err != nil {
			return nil, "", err
		}
		restExp.Children = append(restExp.Children, *operatorArit)

		// <expresion_arit>
		exp, exp_value, err := p.ParseExpression()
		if err != nil {
			return nil, "", err
		}
		restExp.Children = append(restExp.Children, *exp)

		// <rest_arit>
		rest, restExp_value, err := p.ParseRestExp()
		if err != nil {
			return nil, "", err
		}
		restExp.Children = append(restExp.Children, *rest)
		result := operatorArit_value + " " + exp_value + " " + restExp_value
		restExp.TokenValue = result
	} else {
		p.unscan()
	}

	return restExp, restExp.TokenValue, nil
}
