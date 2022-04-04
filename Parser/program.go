package Parser

import (
	"AnalisisLexico/Lexer"
	"fmt"
)

// <AST> -> begin <declarations> <statements> end

// AST is the root of the AST

type AST struct {
	Root     *ASTNode
	Children []ASTNode
}

type ASTNode struct {
	TokenType  Lexer.Token
	TokenValue string
	Children   []ASTNode
}

func (p *Parser) ParseProgram() (*AST, error) {
	AST := &AST{}

	// check for begin
	tok, lit := p.scanIgnoreWhitespace()
	if tok != Lexer.BEGIN {
		return nil, fmt.Errorf("expected begin, got %s", lit)
	}
	// create new node
	// it has no children since it's a terminal node
	AST.Children = append(AST.Children, ASTNode{TokenType: tok, TokenValue: lit})

	// check for declarations
	declaration, d, err := p.ParseDeclarations()
	if err != nil {
		return nil, err
	}

	// add child to the AST
	AST.Children = append(AST.Children, *declaration)

	// check for statement
	statement, s, err := p.ParseStatements()
	if err != nil {
		return nil, err
	}
	AST.Children = append(AST.Children, *statement)

	// check for end
	tok, lit = p.scanIgnoreWhitespace()
	if tok != Lexer.END {
		return nil, fmt.Errorf("expected end, got %s", lit)
	}
	AST.Children = append(AST.Children, ASTNode{TokenType: tok, TokenValue: lit})
	result := fmt.Sprintf("%s %s %s %s", "begin", d, s, "end")
	AST.Root = &ASTNode{TokenType: Lexer.PROGRAM, TokenValue: result}

	return AST, nil
}
