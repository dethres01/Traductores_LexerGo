package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <program> -> begin <declarations> <statements> end

type Program struct {
	abstractSyntaxTree []interface{}
}

func (p *Parser) ParseProgram() (*Program, error) {
	program := &Program{}

	// check for begin
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.BEGIN {
		return nil, fmt.Errorf("expected begin, got %s", lit)
	}
	program.abstractSyntaxTree = append(program.abstractSyntaxTree, lit)

	// check for declarations
	declaration, err := p.ParseDeclarations()
	if err != nil {
		return nil, err
	}
	fmt.Println("declarations: ", declaration)
	program.abstractSyntaxTree = append(program.abstractSyntaxTree, declaration)

	// check for statement
	statement, err := p.ParseStatements()
	if err != nil {
		return nil, err
	}
	fmt.Println("statement: ", statement)
	program.abstractSyntaxTree = append(program.abstractSyntaxTree, statement)

	// check for end
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.END {
		return nil, fmt.Errorf("expected end, got %s", lit)
	}
	return program, nil
}
