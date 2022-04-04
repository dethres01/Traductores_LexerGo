package Semantic

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
)

func getAssignmentInfo(node Parser.ASTNode, s *SemanticAnalysis) error {
	identifier := node.Children[0]
	// check if the identifier is declared
	if !s.SymbolTable.checkExistence(identifier.TokenValue) {
		return fmt.Errorf("variable %s is not declared", identifier.TokenValue)
	}
	// we have to recurse even deeper
	err := getExpressionInfo(node.Children[2], s, identifier)
	if err != nil {
		return err
	}
	return nil
}

func getExpressionInfo(expression Parser.ASTNode, s *SemanticAnalysis, identifier Parser.ASTNode) error {
	// we are a the expression side of assignment
	switch expression.Children[0].TokenType {
	case Lexer.INT:
		// done
		if !s.SymbolTable.compareType(identifier.TokenValue, Lexer.INT) {
			return fmt.Errorf("variable %s is not of type int", identifier.TokenValue)
		}

	case Lexer.FLOAT:
		// done
		if !s.SymbolTable.compareType(identifier.TokenValue, Lexer.FLOAT) {
			return fmt.Errorf("variable %s is not of type float", identifier.TokenValue)
		}
	case Lexer.IDENTIFIER:
		// check if the identifier is declared
		if !s.SymbolTable.checkExistence(expression.Children[0].TokenValue) {
			return fmt.Errorf("variable %s is not declared", expression.Children[0].TokenValue)
		}
		// we have to confirm if the identifiers are the same type
		if !s.SymbolTable.compareTypes(identifier.TokenValue, expression.Children[0].TokenValue) {
			return fmt.Errorf("variable %s is not of the same type as %s", expression.Children[0].TokenValue, identifier.TokenValue)
		}
	case Lexer.LPAREN:
		// syntax is fine so we have to check for
		//<expresion_arit><operador_arit><expresion_arit>
		// expresion_arit is a recursive call
		err := getExpressionInfo(expression.Children[1], s, identifier)
		if err != nil {
			return err
		}
		// we really don't have to do checks for operator arit atm
		// we will have to do it later to make the actual operations tho.
		// RPN
		// children 6 should be rest_exp
		// rest_exp is a recursive call
		err = getRestExpressionInfo(expression.Children[3], s, identifier)
		if err != nil {
			return err
		}

	}

	// after we obtain the first part we have to check for rest_exp
	//<rest_arit> → <operador_arit><expresion_arit><rest_arit> | epsilon

	return nil
}

func getRestExpressionInfo(expression Parser.ASTNode, s *SemanticAnalysis, identifier Parser.ASTNode) error {
	if expression.TokenValue != "" {
		// we don't check operators
		err := getExpressionInfo(expression.Children[1], s, identifier)
		if err != nil {
			return err
		}
		// we recurse call to get the rest of the expression
		err = getRestExpressionInfo(expression.Children[2], s, identifier)
		if err != nil {
			return err
		}
	}
	return nil
}
