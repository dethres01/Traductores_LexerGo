package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

//<tipo> → entero | real

type Type struct {
	AbstractSyntaxTree []interface{}
}

func (p *Parser) ParseType() (*ASTNode, string, error) {
	type_ := &ASTNode{}
	// these are terminal nodes
	fmt.Println("ParseType")
	// they are both terminals we check for them
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.INT && tok != Lexer.FLOAT {
		return nil, "", fmt.Errorf("expected type (int or float), got %s", lit)
	}
	if tok == Lexer.INT {
		type_.TokenType = Lexer.INT
		type_.TokenValue = "int"
	} else {
		type_.TokenType = Lexer.FLOAT
		type_.TokenValue = "float"
	}

	return type_, type_.TokenValue, nil
}
