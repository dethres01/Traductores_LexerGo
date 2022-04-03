package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <program> -> begin <declarations> <statements> end

// Program is the root of the AST

type Program struct {
	AbstractSyntaxTree []ASTNode
}

type ASTNode struct {
	TokenType  Lexer.Token
	TokenValue string
	Children   []ASTNode
}

func (p *Parser) ParseProgram() (*Program, error) {
	program := &Program{}

	// check for begin
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.BEGIN {
		return nil, fmt.Errorf("expected begin, got %s", lit)
	}
	// create new node
	// it has no children since it's a terminal node
	program.AbstractSyntaxTree = append(program.AbstractSyntaxTree, ASTNode{TokenType: tok, TokenValue: lit})

	// check for declarations
	declaration, _, err := p.ParseDeclarations()
	if err != nil {
		return nil, err
	}
	//fmt.Println("declarations: ", declaration)

	// add child to the program
	program.AbstractSyntaxTree = append(program.AbstractSyntaxTree, *declaration)

	// check for statement
	statement, _, err := p.ParseStatements()
	if err != nil {
		return nil, err
	}
	//fmt.Println("statement: ", statement)
	program.AbstractSyntaxTree = append(program.AbstractSyntaxTree, *statement)

	// check for end
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.END {
		return nil, fmt.Errorf("expected end, got %s", lit)
	}
	program.AbstractSyntaxTree = append(program.AbstractSyntaxTree, ASTNode{TokenType: tok, TokenValue: lit})

	return program, nil
}
