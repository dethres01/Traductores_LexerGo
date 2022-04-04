package Semantic

import (
	"AnalisisLexico/Parser"
	"fmt"
)

func getWhileLoopInfo(node Parser.ASTNode, s *SemanticAnalysis) error {
	// syntaxis is okay, we have to validate the condition mainly
	//children[2] is the condition
	//children[4] is statementinfo

	comparison_node := node.Children[2]
	//printSubTree(comparison_node, 0)
	//printSubTree(comparison_node.Children[2], 0)
	//compare type
	if !s.SymbolTable.compareType(comparison_node.Children[0].TokenValue, comparison_node.Children[2].Children[0].TokenType) {
		return fmt.Errorf("type mismatch in condition")
	}
	body_node := node.Children[4]
	printSubTree(body_node, 0)
	err := s.AnalyzeStatements(body_node)
	if err != nil {
		return err
	}
	return nil
}
