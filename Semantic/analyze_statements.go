package Semantic

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
)

// :(

func (s *SemanticAnalysis) AnalyzeStatements(statements Parser.ASTNode) error {
	//check grammar
	//<ordenes> → <orden> **; <rest_ordenes>**
	// check tree
	//printSubTree(statements, 0)
	// works more or less the same as AnalyzeDeclarations
	// but it gets a little complicated becuase this side of the grammar

	// I have to think what do I want do get from the semantic analysis

	// from assign:
	// <asignacion> → <identificador> = <expresion>
	// this is the easier one, I just need to check if the variable is declared, and if the expression is valid
	statement_node := statements.Children[0]
	err := getStatementInfo(statement_node, s)
	if err != nil {
		return err
	}

	rest_statement_node := statements.Children[2]
	if rest_statement_node.TokenValue != "" {
		err := s.AnalyzeStatements(rest_statement_node)
		if err != nil {
			return err
		}
	}

	return nil
}
func getStatementInfo(statement Parser.ASTNode, s *SemanticAnalysis) error {
	//printSubTree(statement, 0)
	// we have to get the children so we can get the proper nodetype
	node := statement.Children[0]
	switch node.TokenType {
	case Lexer.ASSIGNMENT:
		err := getAssignmentInfo(node, s)
		if err != nil {
			return err
		}
		//<rest_arit> → <operador_arit><expresion_arit><rest_arit> | epsilon

	case Lexer.WHILE_LOOP:
		err := getWhileLoopInfo(node, s)
		if err != nil {
			return err
		}

	case Lexer.CONDITION:
		err := getConditionInfo(node, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func printSubTree(node Parser.ASTNode, level int) {
	/*for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("%s\n", node.TokenValue)
	for _, child := range node.Children {
		printSubTree(child, level+1)
	}*/
	for _, child := range node.Children {
		fmt.Println(child.TokenValue)
	}
}
