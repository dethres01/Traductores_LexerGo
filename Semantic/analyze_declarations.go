package Semantic

import (
	"AnalisisLexico/Lexer"
	"AnalisisLexico/Parser"
	"fmt"
)

func (s *SemanticAnalysis) AnalyzeDeclarations(declaration Parser.ASTNode) error {

	tok, variables, err := GetDeclarationInfo(declaration.Children[0])
	if err != nil {
		return err
	}

	for _, variable := range variables {
		err = s.SymbolTable.AddVariable(variable, tok)
		if err != nil {
			return err
		}
	}

	rest_declaration_node := declaration.Children[2]

	if rest_declaration_node.TokenValue != "" {
		// recursion
		err := s.AnalyzeDeclarations(rest_declaration_node)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDeclarationInfo(declaration Parser.ASTNode) (Lexer.Token, []string, error) {
	if len(declaration.Children) != 2 {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected 2 children, got %d", len(declaration.Children))
	}

	type_node := declaration.Children[0]
	if type_node.TokenType != Lexer.INT && type_node.TokenType != Lexer.FLOAT {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected type (int or float) INVALID DATA TYPE ERROR, got %s", type_node.TokenValue)
	}

	dataType := type_node.TokenType

	variables_node := declaration.Children[1]

	// 2.- <rest_lista_variables>
	if len(variables_node.Children) != 2 {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected 2 children, got %d", len(variables_node.Children))
	}

	variable_node := variables_node.Children[0]
	if variable_node.TokenType != Lexer.IDENTIFIER {
		return Lexer.ILLEGAL, []string{}, fmt.Errorf("expected variable, got %s", variable_node.TokenValue)
	}

	variables := append([]string{}, variable_node.TokenValue)

	rest_lista_variables_node := variables_node.Children[1]
	rest_variables, err := searchRestVariables(rest_lista_variables_node)
	if err != nil {
		return Lexer.ILLEGAL, []string{}, err
	}

	variables = append(variables, rest_variables...)

	return dataType, variables, nil
}
func searchVariableList(variables_node Parser.ASTNode) ([]string, error) {
	if len(variables_node.Children) != 2 {
		return []string{}, fmt.Errorf("expected 2 children, got %d", len(variables_node.Children))
	}
	// now we check the first child
	// this child should be a variable
	variable_node := variables_node.Children[0]
	if variable_node.TokenType != Lexer.IDENTIFIER {
		return []string{}, fmt.Errorf("expected variable, got %s", variable_node.TokenValue)
	}
	variables := append([]string{}, variable_node.TokenValue)

	rest_variable_nodes := variables_node.Children[1]
	rest_variables, err := searchRestVariables(rest_variable_nodes)
	if err != nil {
		return []string{}, err
	}
	variables = append(variables, rest_variables...)
	return variables, nil
}

func searchRestVariables(rest_lista_variables_node Parser.ASTNode) ([]string, error) {

	if len(rest_lista_variables_node.Children) != 0 {

		if len(rest_lista_variables_node.Children) != 2 {
			return []string{}, fmt.Errorf("expected 2 children, got %d", len(rest_lista_variables_node.Children))
		}

		comma_node := rest_lista_variables_node.Children[0]
		if comma_node.TokenType != Lexer.COMMA {
			return []string{}, fmt.Errorf("expected comma, got %s", comma_node.TokenValue)
		}

		variables_node := rest_lista_variables_node.Children[1]

		result_variables, err := searchVariableList(variables_node)
		if err != nil {
			return []string{}, err
		}
		return result_variables, nil
	}

	return []string{}, nil
}
