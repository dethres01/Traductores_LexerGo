package Semantic

import "AnalisisLexico/Parser"

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
}
