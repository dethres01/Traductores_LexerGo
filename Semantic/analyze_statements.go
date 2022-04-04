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
	//rest_statement_node := statements.Children[2]

	return nil
}
func getStatementInfo(statement Parser.ASTNode, s *SemanticAnalysis) error {
	//printSubTree(statement, 0)
	// we have to get the children so we can get the proper nodetype
	node := statement.Children[0]
	switch node.TokenType {
	case Lexer.ASSIGNMENT:
		// <asignacion> → <identificador> = <expresion>
		// let's check the tree
		//printSubTree(node, 0)
		// supossed to have 3 children
		if len(node.Children) != 3 {
			return fmt.Errorf("(False Positive on ASSIGN ParserSide), got %d", len(node.Children))
		}
		// <identificador>
		identifier := node.Children[0]
		// check if the identifier is declared
		if !s.SymbolTable.checkExistence(identifier.TokenValue) {
			return fmt.Errorf("variable %s is not declared", identifier.TokenValue)
		}
		// we have confirmed that the identifier is declared

		// so now we have to check the expression
		// <expresion>
		expression := node.Children[2]
		printSubTree(expression, 0)
		// check grammar
		////<expresion_arit> → (<expresion_arit><operador_arit><expresion_arit>) <rest_arit>
		//|	 <identificador> <rest_arit>
		//| <numeros><rest_arit>

		// so we have to check the expression and see which type of expression it is
		// we check for first children tokentype
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

		}
		//<rest_arit> → <operador_arit><expresion_arit><rest_arit> | epsilon

	case Lexer.WHILE_LOOP:

	case Lexer.CONDITION:
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
