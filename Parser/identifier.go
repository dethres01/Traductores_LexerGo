package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

/*
<identificador> → <letra> | <letra><resto_letras>
<letra> → A..Za..z
<resto_letras> → <letraN> | <letraN><resto_letras>
<letraN> → 0..9A..Za..z
*/
// we  really don't check this here since it comes from the lexer

func (p *Parser) ParseIdentifier() (*ASTNode, string, error) {
	identifier := &ASTNode{TokenType: Lexer.IDENTIFIER}
	// these are terminal nodes

	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.ID {
		return nil, "", fmt.Errorf("expected identifier, got %s", lit)
	}
	identifier.TokenValue = lit
	identifier.Children = append(identifier.Children, ASTNode{TokenType: tok, TokenValue: lit})

	return identifier, identifier.TokenValue, nil
}
