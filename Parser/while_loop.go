package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <bucle_while> → while (<comparacion>) <ordenes> endwhile

func (p *Parser) ParseWhileLoop() (*ASTNode, string, error) {
	whileLoop := &ASTNode{TokenType: Lexer.WHILE_LOOP}

	// check for while
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.WHILE {
		return nil, "", fmt.Errorf("expected while, got %s", lit)
	}
	whileLoop.Children = append(whileLoop.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for (
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.LPAREN {
		return nil, "", fmt.Errorf("expected (, got %s", lit)
	}
	whileLoop.Children = append(whileLoop.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for <comparacion>
	comparacion, comparison_value, err := p.ParseComparison()
	if err != nil {
		return nil, "", err
	}
	whileLoop.Children = append(whileLoop.Children, *comparacion)
	p.ic.while(comparison_value)
	// check for )
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.RPAREN {
		return nil, "", fmt.Errorf("expected ), got %s", lit)
	}
	whileLoop.Children = append(whileLoop.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for <ordenes>
	ordenes, statements_value, err := p.ParseStatements()
	if err != nil {
		return nil, "", err
	}
	whileLoop.Children = append(whileLoop.Children, *ordenes)

	// check for endwhile
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.ENDWHILE {
		return nil, "", fmt.Errorf("expected endwhile, got %s", lit)
	}
	p.ic.EndWhile(comparison_value)
	whileLoop.Children = append(whileLoop.Children, ASTNode{TokenType: tok, TokenValue: lit})
	result := fmt.Sprintf("while (%s) %s endwhile", comparison_value, statements_value)
	whileLoop.TokenValue = result
	return whileLoop, whileLoop.TokenValue, nil
}
