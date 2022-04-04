package Semantic

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
)

func getConditionInfo(node Parser.ASTNode, s *SemanticAnalysis) error {
	// syntaxis is okay, we have to validate the condition mainly
	//printSubTree(node, 0)
	//children[2] is the condition
	//children[4] is statementinfo
	//children[5] is rest_condition

	comparison_node := node.Children[2]
	if !s.SymbolTable.compareType(comparison_node.Children[0].TokenValue, comparison_node.Children[2].Children[0].TokenType) {
		return fmt.Errorf("type mismatch in condition")
	}
	body_node := node.Children[4]
	err := s.AnalyzeStatements(body_node)
	if err != nil {
		return err
	}
	rest_condition_node := node.Children[5]
	err = getRestConditionInfo(rest_condition_node, s)
	if err != nil {
		return err
	}
	return nil
}

func getRestConditionInfo(node Parser.ASTNode, s *SemanticAnalysis) error {
	if node.Children[0].TokenType == Lexer.ELSE {
		//chidlren[1] is the body
		body_node := node.Children[1]
		err := s.AnalyzeStatements(body_node)
		if err != nil {
			return err
		}
	}
	return nil
}
