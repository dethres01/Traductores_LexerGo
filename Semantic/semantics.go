package Semantic

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
)

//the purpose of this package is to provide a semantic analysis
//for the language.
// we are going to use the parser package to get the AST
// and then we will use the AST to do semantic analysis
// and then we will use the AST to generate the code
// and then we will use the code to generate the bytecode
// and then we will use the bytecode to execute the program

//let's get the AST to make it a variable in this package
type SemanticAnalysis struct {
	AST *Parser.AST
	//SymbolTable *SymbolTable
	SymbolTable *SymbolTable
}

//symbol table
type SymbolTable struct {
	SymbolTable map[string]*Symbol
}

//Symbol
type Symbol struct {
	//name of the symbol
	Name string
	//type of the symbol
	Type Lexer.Token
	//value of the symbol, if it is a variable
	Value interface{}
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		SymbolTable: make(map[string]*Symbol),
	}
}

func NewSymbol(name string, type_ Lexer.Token, value interface{}) *Symbol {
	return &Symbol{
		Name:  name,
		Type:  type_,
		Value: value,
	}
}
func (sb *SymbolTable) AddVariable(name string, type_ Lexer.Token) error {
	//we need to check if the variable is already declared
	if _, ok := sb.SymbolTable[name]; ok {
		// if it is already declared we need to throw an error
		// we can do this by returning an error
		return fmt.Errorf("variable %s already declared", name)
	}
	sb.SymbolTable[name] = NewSymbol(name, type_, nil)
	return nil
}
func NewSemanticAnalysis(ast *Parser.AST) *SemanticAnalysis {
	return &SemanticAnalysis{
		AST: ast,
	}
}

func (s *SemanticAnalysis) Start() error {
	//we need to start by checking the AST

	// initialize symbol table
	s.SymbolTable = NewSymbolTable()

	// we checked the syntax so we don't need to check it again
	// but we can check if the AST is empty
	if s.AST.Root == nil {
		return fmt.Errorf("AST is empty")
	}
	// we can start by checking the root node
	// we need to check if it is a program
	if s.AST.Root.TokenType != Lexer.PROGRAM {
		return fmt.Errorf("expected program, got %s", s.AST.Root.TokenValue)
	}
	// we need to check if the program has a block
	if len(s.AST.Children) != 4 {
		return fmt.Errorf("expected 4 Children, got %d", len(s.AST.Children))
	}
	// we need to check if the block has a begin on the first position
	if s.AST.Children[0].TokenType != Lexer.BEGIN {
		return fmt.Errorf("expected begin, got %s", s.AST.Root.Children[0].TokenValue)
	}
	// we need to check if the block has an end on the last position
	if s.AST.Children[3].TokenType != Lexer.END {
		return fmt.Errorf("expected end, got %s", s.AST.Root.Children[3].TokenValue)
	}
	// we need to check if the block has declarations on the second position
	if s.AST.Children[1].TokenType != Lexer.DECLARATIONS {
		return fmt.Errorf("expected declarations, got %s", s.AST.Root.Children[1].TokenValue)
	}
	// we need to check if the block has statements on the third position
	if s.AST.Children[2].TokenType != Lexer.STATEMENTS {
		return fmt.Errorf("expected statements, got %s", s.AST.Root.Children[2].TokenValue)
	}

	//now we validated the general structure of the AST

	//we should start checking if the declarations are correct
	// because we need to check if the variables are declared

	err := s.AnalyzeDeclarations(s.AST.Children[1])
	if err != nil {
		return err
	}

	return nil
}
