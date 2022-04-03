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

type Identifier struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseIdentifier() (*Identifier, error) {
	Identifier := &Identifier{}

	fmt.Println("ParseIdentifier")

	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.ID {
		return nil, fmt.Errorf("expected identifier, got %s", lit)
	}
	Identifier.AbstractSyntaxTree = append(Identifier.AbstractSyntaxTree, lit)

	return Identifier, nil
}