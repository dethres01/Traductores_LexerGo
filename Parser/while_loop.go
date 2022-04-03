package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <bucle_while> → while (<comparacion>) <ordenes> endwhile

type WhileLoop struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseWhileLoop() (*WhileLoop, error) {
	whileLoop := &WhileLoop{}

	fmt.Println("ParseWhileLoop")
	// check for while
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.WHILE {
		return nil, fmt.Errorf("expected while, got %s", lit)
	}
	whileLoop.AbstractSyntaxTree = append(whileLoop.AbstractSyntaxTree, lit)

	// check for (
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.LPAREN {
		return nil, fmt.Errorf("expected (, got %s", lit)
	}
	whileLoop.AbstractSyntaxTree = append(whileLoop.AbstractSyntaxTree, lit)

	// check for <comparacion>
	comparacion, err := p.ParseComparison()
	if err != nil {
		return nil, err
	}
	whileLoop.AbstractSyntaxTree = append(whileLoop.AbstractSyntaxTree, comparacion)

	// check for )
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.RPAREN {
		return nil, fmt.Errorf("expected ), got %s", lit)
	}
	whileLoop.AbstractSyntaxTree = append(whileLoop.AbstractSyntaxTree, lit)

	// check for <ordenes>
	ordenes, err := p.ParseStatements()
	if err != nil {
		return nil, err
	}
	whileLoop.AbstractSyntaxTree = append(whileLoop.AbstractSyntaxTree, ordenes)

	// check for endwhile
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.ENDWHILE {
		return nil, fmt.Errorf("expected endwhile, got %s", lit)
	}
	whileLoop.AbstractSyntaxTree = append(whileLoop.AbstractSyntaxTree, lit)

	return whileLoop, nil
}
